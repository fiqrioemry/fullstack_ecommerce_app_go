package services

import (
	"errors"
	"fmt"
	"log"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"

	"gorm.io/gorm"
)

type PaymentService interface {
	ExpireOldPendingPayments() error
	HandlePaymentNotification(req dto.MidtransNotificationRequest) error
	GetAllUserPayments(param dto.PaymentQueryParam) ([]dto.AdminPaymentResponse, *dto.PaginationResponse, error)
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
		return err
	}

	if payment.Status == "success" {
		return nil
	}

	payment.Method = req.PaymentType

	switch req.TransactionStatus {
	case "settlement", "capture":
		if req.FraudStatus == "accept" || req.FraudStatus == "" {
			payment.Status = "success"
		}
	case "pending":
		payment.Status = "pending"
	default:
		payment.Status = "failed"

		if err := s.productRepo.RestoreStockOnPaymentFailure(&payment.Order); err != nil {
			return fmt.Errorf("restore stock failed: %w", err)
		}
		if err := s.orderRepo.UpdateOrder(&models.Order{
			ID:     payment.Order.ID,
			Status: "canceled",
		}); err != nil {
			return err
		}
	}

	if err := s.paymentRepo.UpdatePayment(payment); err != nil {
		return err
	}

	if payment.Status == "success" {

		// TODO: Replace with RabbitMQ for async notification dispatch ---
		// ? Send notification : success payment info
		payload := dto.NotificationEvent{
			UserID:  payment.UserID.String(),
			Type:    "order_processed",
			Title:   "Payment Successfully Received",
			Message: "Thank you for your payment. Your order is now being processed.",
		}

		err = s.notificationService.SendToUser(payload)
		if err != nil {
			log.Printf("Gagal mengirim notifikasi ke pengguna %s: %v\n", payload.UserID, err)
		}
		// TODO: Replace with RabbitMQ for async notification dispatch ---

		if err := s.orderRepo.UpdateOrder(&models.Order{
			ID:     payment.Order.ID,
			Status: "pending",
		}); err != nil {
			return err
		}

		if payment.Order.VoucherCode != nil && *payment.Order.VoucherCode != "" {
			if err := s.voucherService.DecreaseQuota(payment.UserID, *payment.Order.VoucherCode); err != nil {
				return err
			}
		}

		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

	}

	return nil
}

func (s *paymentService) GetAllUserPayments(param dto.PaymentQueryParam) ([]dto.AdminPaymentResponse, *dto.PaginationResponse, error) {
	payments, total, err := s.paymentRepo.GetAllUserPayments(param)
	if err != nil {
		return nil, nil, err
	}

	var result []dto.AdminPaymentResponse
	for _, p := range payments {
		result = append(result, dto.AdminPaymentResponse{
			ID:            p.ID.String(),
			UserID:        p.UserID.String(),
			OrderID:       p.OrderID.String(),
			InvoiceNumber: p.Order.InvoiceNumber,
			UserEmail:     p.User.Email,
			Fullname:      p.User.Profile.Fullname,
			Total:         p.Total,
			Method:        p.Method,
			Status:        p.Status,
			PaidAt:        p.PaidAt.Format("2006-01-02 15:04:05"),
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

		if err := s.paymentRepo.UpdatePayment(&p); err != nil {
			return fmt.Errorf("failed to update payment %s: %w", p.ID, err)
		}

		if err := s.productRepo.RestoreStockOnPaymentFailure(&p.Order); err != nil {
			return fmt.Errorf("failed to restore stock for order %s: %w", p.OrderID, err)
		}

		if err := s.orderRepo.UpdateOrder(&models.Order{
			ID:     p.Order.ID,
			Status: "canceled",
		}); err != nil {
			return fmt.Errorf("failed to update order status: %w", err)
		}
	}

	log.Printf("✅ %d payments expired → failed, orders canceled, and stock restored\n", len(payments))
	return nil
}
