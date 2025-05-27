package repositories

import (
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	CreateReview(review *models.Review) error
	MarkItemAsReviewed(itemID uuid.UUID) error
	GetOrderItemByID(itemID string) (*models.OrderItem, error)
	GetReviewsByProductID(productID uuid.UUID) ([]models.Review, error)
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db}
}

func (r *reviewRepository) CreateReview(review *models.Review) error {
	return r.db.Create(review).Error
}

func (r *reviewRepository) MarkItemAsReviewed(itemID uuid.UUID) error {
	return r.db.Model(&models.OrderItem{}).Where("id = ?", itemID).
		Update("is_reviewed", true).Error
}

func (r *reviewRepository) GetOrderItemByID(itemID string) (*models.OrderItem, error) {
	var item models.OrderItem
	err := r.db.Where("id = ?", itemID).First(&item).Error
	if err != nil {
		return nil, err
	}
	return &item, nil
}

func (r *reviewRepository) GetReviewsByProductID(productID uuid.UUID) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("product_id = ?", productID).
		Order("created_at DESC").Preload("User.Profile").Find(&reviews).Error
	return reviews, err
}
