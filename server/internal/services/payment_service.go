package services

import (
	"fmt"
	"log"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"time"
)

type PaymentService interface {
	ExpireOldPendingPayments() error
	HandlePaymentNotification(req dto.MidtransNotificationRequest) error
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
func (s *paymentService) HandlePaymentNotification(req dto.MidtransNotificationRequest) error {
	payment, err := s.paymentRepo.GetPaymentByOrderID(req.OrderID)
	if err != nil {
		return fmt.Errorf("payment not found: %w", err)
	}

	if payment.Status == "success" {
		log.Printf("Payment already marked as success for order %s", req.OrderID)
		return nil
	}

	payment.Method = req.PaymentType

	switch req.TransactionStatus {

	case "settlement", "capture":
		if req.FraudStatus == "accept" || req.FraudStatus == "" {
			payment.Status = "success"
			payment.PaidAt = time.Now()
		} else {
			payment.Status = "failed"
			err = s.productRepo.RestoreStockOnPaymentFailure(&payment.Order)
			if err != nil {
				return fmt.Errorf("failed to restore stock for order %s: %w", payment.OrderID, err)
			}

			err = s.orderRepo.UpdateOrder(&models.Order{
				ID:     payment.Order.ID,
				Status: "canceled",
			})
			if err != nil {
				return fmt.Errorf("failed to update order status: %w", err)
			}
		}
	case "pending":
		payment.Status = "pending"
	default:
		payment.Status = "failed"
		if req.FraudStatus == "accept" || req.FraudStatus == "" {
			payment.Status = "success"
			payment.PaidAt = time.Now()
		} else {
			payment.Status = "failed"
			err = s.productRepo.RestoreStockOnPaymentFailure(&payment.Order)
			if err != nil {
				return fmt.Errorf("failed to restore stock for order %s: %w", payment.OrderID, err)
			}

			err = s.orderRepo.UpdateOrder(&models.Order{
				ID:     payment.Order.ID,
				Status: "canceled",
			})
			if err != nil {
				return fmt.Errorf("failed to update order status: %w", err)
			}
		}
	}

	if err := s.paymentRepo.UpdatePayment(payment); err != nil {
		return fmt.Errorf("failed to update payment: %w", err)
	}

	if payment.Status == "success" {
		if err := s.orderRepo.UpdateOrder(&models.Order{
			ID:     payment.Order.ID,
			Status: "pending",
		}); err != nil {
			return fmt.Errorf("failed to update order: %w", err)
		}

		notification := dto.NotificationEvent{
			UserID: payment.UserID.String(),
			Type:   "order_processed",
			Title:  "Payment Successfully Received",
			Message: fmt.Sprintf("Thank you %s, your payment for order %s has been received and is being processed.",
				payment.Fullname, payment.Order.InvoiceNumber),
		}
		if err := s.notificationService.SendToUser(notification); err != nil {
			log.Printf("Failed sending notification to user %s: %v", notification.UserID, err)
		} else {
			log.Println("Notification sent to user")
		}
	}

	return nil
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
