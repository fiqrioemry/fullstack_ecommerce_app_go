package repositories

import (
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	CreateReview(review *models.Review) error
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

func (r *reviewRepository) GetReviewsByProductID(productID uuid.UUID) ([]models.Review, error) {
	var reviews []models.Review
	err := r.db.Where("product_id = ?", productID).
		Order("created_at DESC").Preload("User.Profile").Find(&reviews).Error
	return reviews, err
}
