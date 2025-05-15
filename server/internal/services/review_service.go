package services

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"

	"github.com/google/uuid"
)

type ReviewService interface {
	CreateReview(orderID, userID, productID string, req dto.CreateReviewRequest) error
	GetReviewsByProductID(productID string) ([]dto.ReviewResponse, error)
}

type reviewService struct {
	reviewRepo repositories.ReviewRepository
	orderRepo  repositories.OrderRepository
}

func NewReviewService(reviewRepo repositories.ReviewRepository, orderRepo repositories.OrderRepository) ReviewService {
	return &reviewService{reviewRepo, orderRepo}
}

func (s *reviewService) CreateReview(orderID, userID, productID string, req dto.CreateReviewRequest) error {
	uid, _ := uuid.Parse(userID)
	pid, _ := uuid.Parse(productID)

	order, err := s.orderRepo.GetOrderDetail(orderID)
	if err != nil || order.UserID != uid {
		return errors.New("unauthorized or order not found")
	}
	if order.Shipment.Status != "delivered" {
		return errors.New("cannot review before product is delivered")
	}

	// Pastikan product ada dalam order
	found := false
	for _, item := range order.Items {
		if item.ProductID == pid {
			found = true
			break
		}
	}
	if !found {
		return errors.New("product not found in this order")
	}

	review := &models.Review{
		ID:        uuid.New(),
		UserID:    uid,
		ProductID: pid,
		Rating:    req.Rating,
		Comment:   req.Comment,
		ImageURL:  req.ImageURL,
	}
	return s.reviewRepo.CreateReview(review)
}

func (s *reviewService) GetReviewsByProductID(productID string) ([]dto.ReviewResponse, error) {
	pid, err := uuid.Parse(productID)
	if err != nil {
		return nil, errors.New("invalid product ID")
	}
	reviews, err := s.reviewRepo.GetReviewsByProductID(pid)
	if err != nil {
		return nil, err
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
			ImageURL:  r.ImageURL,
			CreatedAt: r.CreatedAt,
		})
	}
	return result, nil
}
