package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v75"
)

type PaymentService interface {
	ExpireOldPendingPayments() error
	WebhookNotifications(c *gin.Context)
	GetAllUserPayments(param dto.PaymentQueryParam) ([]dto.PaymentResponse, *dto.PaginationResponse, error)
}

type paymentService struct {
	paymentRepo         repositories.PaymentRepository
	authRepo            repositories.AuthRepository
	productRepo         repositories.ProductRepository
	voucherService      VoucherService
	orderRepo           repositories.OrderRepository
	notificationService NotificationService
}

func NewPaymentService(
	paymentRepo repositories.PaymentRepository,
	authRepo repositories.AuthRepository,
	productRepo repositories.ProductRepository,
	voucherService VoucherService,
	orderRepo repositories.OrderRepository,
	notificationService NotificationService,
) PaymentService {
	return &paymentService{
		paymentRepo:         paymentRepo,
		authRepo:            authRepo,
		productRepo:         productRepo,
		voucherService:      voucherService,
		orderRepo:           orderRepo,
		notificationService: notificationService,
	}
}

func (s *paymentService) WebhookNotifications(c *gin.Context) {
	const MaxBodyBytes = int64(65536)
	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)

	body, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read request body"})
		return
	}

	event := stripe.Event{}
	if err := json.Unmarshal(body, &event); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}

	switch event.Type {
	case "checkout.session.completed":
		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session data"})
			return
		}

		orderID := session.Metadata["order_id"]
		payment, err := s.paymentRepo.GetPaymentByOrderID(orderID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if payment.Status == "success" {
			c.JSON(http.StatusOK, gin.H{"message": "payment already processed"})
			return
		}

		payment.Method = "card"
		payment.Status = "success"
		if err := s.paymentRepo.UpdatePayment(payment); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err := s.orderRepo.UpdateOrder(&models.Order{
			ID:     payment.Order.ID,
			Status: "pending",
		}); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// ? Waiting for order is process notifications : event 2
		// TODO : Replace with rabbit mq for sending async event ----
		notification := dto.NotificationEvent{
			UserID:  payment.UserID.String(),
			Type:    "order_processed",
			Title:   "Payment Successfully Received",
			Message: fmt.Sprintf("Thank you for %s your payment. Your order is now being processed by admin.", payment.Fullname),
		}
		if err := s.notificationService.SendToUser(notification); err != nil {
			log.Printf("failed sending notification to user %s: %v\n", notification.UserID, err)
		}
		// TODO : Replace with rabbit mq for sending async event ----
	default:
		log.Printf("Unhandled event type: %s\n", event.Type)
	}

	c.JSON(http.StatusOK, gin.H{"message": "payment successfully updated"})
}

func (s *paymentService) GetAllUserPayments(param dto.PaymentQueryParam) ([]dto.PaymentResponse, *dto.PaginationResponse, error) {
	payments, total, err := s.paymentRepo.GetAllUserPayments(param)
	if err != nil {
		return nil, nil, err
	}

	var result []dto.PaymentResponse
	for _, p := range payments {
		result = append(result, dto.PaymentResponse{
			ID:            p.ID.String(),
			UserID:        p.UserID.String(),
			OrderID:       p.OrderID.String(),
			InvoiceNumber: p.Order.InvoiceNumber,
			UserEmail:     p.Email,
			Fullname:      p.Fullname,
			Total:         p.Total,
			Method:        p.Method,
			Status:        p.Status,
			PaidAt:        p.PaidAt.Format("2006-01-02"),
		})
	}
	totalPages := int((total + int64(param.Limit) - 1) / int64(param.Limit))
	pagination := &dto.PaginationResponse{
		Page:       param.Page,
		Limit:      param.Limit,
		TotalRows:  int(total),
		TotalPages: totalPages,
	}

	return result, pagination, nil
}

func (s *paymentService) ExpireOldPendingPayments() error {
	payments, err := s.paymentRepo.GetExpiredPendingPayments()
	if err != nil {
		return fmt.Errorf("failed to fetch expired payments: %w", err)
	}

	if len(payments) == 0 {
		log.Println("✅ No expired pending payments found")
		return nil
	}

	for _, p := range payments {
		p.Status = "failed"

		err := s.paymentRepo.UpdatePayment(&p)
		if err != nil {
			return fmt.Errorf("failed to update payment %s: %w", p.ID, err)
		}

		err = s.productRepo.RestoreStockOnPaymentFailure(&p.Order)
		if err != nil {
			return fmt.Errorf("failed to restore stock for order %s: %w", p.OrderID, err)
		}

		err = s.orderRepo.UpdateOrder(&models.Order{
			ID:     p.Order.ID,
			Status: "canceled",
		})
		if err != nil {
			return fmt.Errorf("failed to update order status: %w", err)
		}
	}

	log.Printf("✅ %d payments expired → failed, orders canceled, and stock restored\n", len(payments))
	return nil
}
