package services

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"

	"github.com/google/uuid"
)

type CartService interface {
	ClearCart(userID string) error
	RemoveItem(userID, productID string) error
	AddToCart(userID string, req dto.CartItemRequest) error
	UpdateQuantity(userID, productID string, quantity int) error
	GetCart(userID string) ([]dto.CartItemResponse, float64, error)
	ToggleItemChecked(userID, productID string) error
}

type cartService struct {
	cartRepo    repositories.CartRepository
	productRepo repositories.ProductRepository
}

func NewCartService(cartRepo repositories.CartRepository, productRepo repositories.ProductRepository) CartService {
	return &cartService{cartRepo, productRepo}
}

func (s *cartService) GetCart(userID string) ([]dto.CartItemResponse, float64, error) {
	uid, _ := uuid.Parse(userID)
	carts, err := s.cartRepo.GetByUserID(uid)
	if err != nil {
		return nil, 0, err
	}

	var items []dto.CartItemResponse
	var total float64

	for _, c := range carts {
		if len(c.Product.ProductGallery) == 0 {
			continue
		}

		price := c.Product.Price
		discount := 0.0
		if c.Product.Discount != nil {
			discount = *c.Product.Discount
		}

		discountedPrice := price * (1 - discount/100)
		originalSubtotal := price * float64(c.Quantity)
		discountedSubtotal := discountedPrice * float64(c.Quantity)

		total += discountedSubtotal

		items = append(items, dto.CartItemResponse{
			ProductID:        c.ProductID.String(),
			Name:             c.Product.Name,
			Price:            price,
			Discount:         discount,
			DiscountedPrice:  discountedPrice,
			Image:            c.Product.ProductGallery[0].Image,
			IsChecked:        c.IsChecked,
			Weight:           c.Product.Weight,
			Quantity:         c.Quantity,
			OriginalSubtotal: originalSubtotal,
			Subtotal:         discountedSubtotal,
		})
	}

	return items, total, nil
}

func (s *cartService) AddToCart(userID string, req dto.CartItemRequest) error {
	uid, _ := uuid.Parse(userID)
	pid, _ := uuid.Parse(req.ProductID)

	product, err := s.productRepo.GetProductByID(pid)
	if err != nil {
		return err
	}
	if req.Quantity > product.Stock {
		return errors.New("stock not available")
	}

	cart := &models.Cart{
		UserID:    uid,
		ProductID: pid,
		Quantity:  req.Quantity,
	}

	return s.cartRepo.AddOrUpdate(cart)
}

func (s *cartService) UpdateQuantity(userID, productID string, quantity int) error {
	uid, _ := uuid.Parse(userID)
	pid, _ := uuid.Parse(productID)

	product, err := s.productRepo.GetProductByID(pid)
	if err != nil {
		return err
	}
	if quantity > product.Stock {
		return errors.New("stock not available")
	}

	return s.cartRepo.UpdateQuantity(uid, pid, quantity)
}

func (s *cartService) RemoveItem(userID, productID string) error {
	uid, _ := uuid.Parse(userID)
	pid, _ := uuid.Parse(productID)
	return s.cartRepo.RemoveItem(uid, pid)
}

func (s *cartService) ClearCart(userID string) error {
	uid, _ := uuid.Parse(userID)
	return s.cartRepo.Clear(uid)
}

func (s *cartService) ToggleItemChecked(userID, productID string) error {
	uid, _ := uuid.Parse(userID)
	pid, _ := uuid.Parse(productID)
	return s.cartRepo.ToggleIsChecked(uid, pid)
}
