package repositories

import (
	"server/internal/dto"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	CreateReview(review *models.Review) error
	MarkItemAsReviewed(itemID uuid.UUID) error
	GetOrderItemByID(itemID string) (*models.OrderItem, error)
	GetReviewsByProductID(productID uuid.UUID, param dto.ReviewQueryParam) ([]models.Review, int64, error)
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

func (r *reviewRepository) GetReviewsByProductID(productID uuid.UUID, param dto.ReviewQueryParam) ([]models.Review, int64, error) {
	var reviews []models.Review
	var count int64

	page := param.Page
	if page <= 0 {
		page = 1
	}
	limit := param.Limit
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	db := r.db.Model(&models.Review{}).
		Where("product_id = ?", productID).
		Preload("User.Profile").
		Order("created_at desc")

	if err := db.Count(&count).Error; err != nil {
		return nil, 0, err
	}

	if err := db.Offset(offset).Limit(limit).Find(&reviews).Error; err != nil {
		return nil, 0, err
	}

	return reviews, count, nil
}
