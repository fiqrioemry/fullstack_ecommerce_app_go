package repositories

import (
	"server/internal/dto"
	"server/internal/models"
	"time"

	"gorm.io/gorm"
)

type PaymentRepository interface {
	CreatePayment(payment *models.Payment) error
	UpdatePayment(payment *models.Payment) error
	GetPaymentByID(id string) (*models.Payment, error)
	GetExpiredPendingPayments() ([]models.Payment, error)
	GetPaymentByOrderID(orderID string) (*models.Payment, error)
	GetAllUserPayments(param dto.PaymentQueryParam) ([]models.Payment, int64, error)
}

type paymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) CreatePayment(payment *models.Payment) error {
	return r.db.Create(payment).Error
}

func (r *paymentRepository) GetPaymentByID(id string) (*models.Payment, error) {
	var payment models.Payment
	if err := r.db.First(&payment, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) GetPaymentByOrderID(orderID string) (*models.Payment, error) {
	var payment models.Payment
	if err := r.db.
		Preload("Order").
		First(&payment, "order_id = ?", orderID).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentRepository) UpdatePayment(payment *models.Payment) error {
	return r.db.Save(payment).Error
}

func (r *paymentRepository) GetAllUserPayments(param dto.PaymentQueryParam) ([]models.Payment, int64, error) {
	var payments []models.Payment
	var count int64

	// Validasi pagination
	page := param.Page
	if page <= 0 {
		page = 1
	}
	limit := param.Limit
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	// Base query
	db := r.db.Model(&models.Payment{}).
		Preload("Order", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		})

	// Search
	if param.Search != "" {
		search := "%" + param.Search + "%"
		db = db.Where("email LIKE ? OR fullname LIKE ?", search, search)
	}

	// Filter status
	if param.Status != "" && param.Status != "all" {
		db = db.Where("payments.status = ?", param.Status)
	}

	// Sorting
	sort := "paid_at desc"
	switch param.Sort {
	case "paid_at_asc":
		sort = "paid_at asc"
	case "paid_at_desc":
		sort = "paid_at desc"
	case "total_asc":
		sort = "total asc"
	case "total_desc":
		sort = "total desc"
	}
	db = db.Order(sort)

	//  total data
	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// dengan pagination
	if err := db.Offset(offset).Limit(limit).Find(&payments).Error; err != nil {
		return nil, 0, err
	}

	return payments, count, nil
}

// ** khusus cron job update status ke failed
func (r *paymentRepository) GetExpiredPendingPayments() ([]models.Payment, error) {
	var payments []models.Payment
	threshold := time.Now().Add(-24 * time.Hour)

	err := r.db.
		Preload("Order.Items").
		Where("status = ? AND paid_at <= ?", "waiting_payment", threshold).
		Find(&payments).Error

	return payments, err
}
