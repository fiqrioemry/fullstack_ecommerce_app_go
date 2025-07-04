package services

import (
	"errors"
	"fmt"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"

	"github.com/google/uuid"
)

type ReviewService interface {
	CreateReview(userID, productID string, req dto.CreateReviewRequest) error
	GetReviewsByProductID(productID string, param dto.ReviewQueryParam) ([]dto.ReviewResponse, *dto.PaginationResponse, error)
}

type reviewService struct {
	reviewRepo repositories.ReviewRepository
	orderRepo  repositories.OrderRepository
}

func NewReviewService(reviewRepo repositories.ReviewRepository, orderRepo repositories.OrderRepository) ReviewService {
	return &reviewService{reviewRepo, orderRepo}
}

func (s *reviewService) CreateReview(userID, itemID string, req dto.CreateReviewRequest) error {
	uid, _ := uuid.Parse(userID)

	item, err := s.reviewRepo.GetOrderItemByID(itemID)
	if err != nil {
		return fmt.Errorf("item not found: %v", err)
	}

	review := &models.Review{
		ID:        uuid.New(),
		UserID:    uid,
		ProductID: item.ProductID,
		Rating:    req.Rating,
		Comment:   req.Comment,
		Image:     &req.ImageURL,
	}

	if err := s.reviewRepo.CreateReview(review); err != nil {
		return fmt.Errorf("failed to create review: %v", err)
	}

	if err := s.reviewRepo.MarkItemAsReviewed(item.ID); err != nil {
		return fmt.Errorf("failed to update OrderItem review status: %v", err)
	}

	return nil
}


func (s *reviewService) GetReviewsByProductID(productID string, param dto.ReviewQueryParam) ([]dto.ReviewResponse, *dto.PaginationResponse, error) {
	pid, err := uuid.Parse(productID)
	if err != nil {
		return nil, nil, errors.New("invalid product ID")
	}

	reviews, total, err := s.reviewRepo.GetReviewsByProductID(pid, param)
	if err != nil {
		return nil, nil, err
	}

	var result []dto.ReviewResponse
	for _, r := range reviews {
		result = append(result, dto.ReviewResponse{
			ID:        r.ID.String(),
			UserID:    r.UserID.String(),
			Fullname:  r.User.Profile.Fullname,
			Avatar:    r.User.Profile.Avatar,
			ProductID: r.ProductID.String(),
			Rating:    r.Rating,
			Comment:   r.Comment,
			Image:     r.Image,
			CreatedAt: r.CreatedAt,
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
