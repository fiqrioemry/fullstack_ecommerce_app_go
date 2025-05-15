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
	if err := r.db.First(&payment, "id = ?", orderID).Error; err != nil {
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

	db := r.db.Model(&models.Payment{}).
		Preload("Order", func(db *gorm.DB) *gorm.DB {
			return db.Unscoped()
		}).
		Preload("User.Profile")

	if param.Search != "" {
		search := "%" + param.Search + "%"
		db = db.Joins("JOIN users ON users.id = payments.user_id").
			Where("users.email LIKE ? OR users.id LIKE ?", search, search)
	}

	if param.Status != "" {
		db = db.Where("payments.status = ?", param.Status)
	}

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	sort := "paid_at desc"
	switch param.Sort {
	case "paid_at asc":
		sort = "paid_at asc"
	case "paid_at desc":
		sort = "paid_at desc"
	case "total asc":
		sort = "total asc"
	case "total desc":
		sort = "total desc"
	case "status asc":
		sort = "status asc"
	case "status desc":
		sort = "status desc"
	}
	db = db.Order(sort)

	offset := (param.Page - 1) * param.Limit
	if err := db.Limit(param.Limit).Offset(offset).Find(&payments).Error; err != nil {
		return nil, 0, err
	}

	return payments, count, nil
}

// ** khusus cron job update status ke failed
func (r *paymentRepository) GetExpiredPendingPayments() ([]models.Payment, error) {
	var payments []models.Payment
	threshold := time.Now().Add(-2 * time.Hour)

	err := r.db.
		Preload("Order.Items").
		Where("status = ? AND paid_at <= ?", "pending", threshold).
		Find(&payments).Error

	return payments, err
}
