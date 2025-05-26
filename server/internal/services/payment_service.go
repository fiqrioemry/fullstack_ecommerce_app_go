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
	log.Println("🔔 Webhook endpoint called")

	c.Request.Body = http.MaxBytesReader(c.Writer, c.Request.Body, MaxBodyBytes)

	body, err := c.GetRawData()
	if err != nil {
		log.Println("❌ Failed to read request body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read request body"})
		return
	}
	log.Println("✅ Request body read successfully")

	var event stripe.Event
	if err := json.Unmarshal(body, &event); err != nil {
		log.Println("❌ Failed to unmarshal JSON into event:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid JSON"})
		return
	}
	log.Printf("📦 Event received: %s\n", event.Type)

	switch event.Type {
	case "checkout.session.completed":
		var session stripe.CheckoutSession
		if err := json.Unmarshal(event.Data.Raw, &session); err != nil {
			log.Println("❌ Failed to unmarshal session data:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid session data"})
			return
		}
		log.Printf("🧾 Session parsed. OrderID: %s\n", session.Metadata["order_id"])

		orderID := session.Metadata["order_id"]
		payment, err := s.paymentRepo.GetPaymentByOrderID(orderID)
		if err != nil {
			log.Printf("❌ Failed to get payment by order ID %s: %v\n", orderID, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Printf("💳 Payment retrieved. Current status: %s\n", payment.Status)

		if payment.Status == "success" {
			log.Println("ℹ️ Payment already marked as success. Skipping update.")
			c.JSON(http.StatusOK, gin.H{"message": "payment already processed"})
			return
		}

		payment.Method = "card"
		payment.Status = "success"
		if err := s.paymentRepo.UpdatePayment(payment); err != nil {
			log.Printf("❌ Failed to update payment: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Println("✅ Payment status updated to success")

		if err := s.orderRepo.UpdateOrder(&models.Order{
			ID:     payment.Order.ID,
			Status: "pending",
		}); err != nil {
			log.Printf("❌ Failed to update order status: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Println("✅ Order status updated to pending")

		notification := dto.NotificationEvent{
			UserID:  payment.UserID.String(),
			Type:    "order_processed",
			Title:   "Payment Successfully Received",
			Message: fmt.Sprintf("Thank you for %s your payment. Your order is now being processed by admin.", payment.Fullname),
		}
		if err := s.notificationService.SendToUser(notification); err != nil {
			log.Printf("❌ Failed sending notification to user %s: %v\n", notification.UserID, err)
		} else {
			log.Println("📨 Notification sent to user")
		}

	default:
		log.Printf("⚠️ Unhandled event type: %s\n", event.Type)
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
