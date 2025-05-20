package repositories

import (
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CartRepository interface {
	Clear(userID uuid.UUID) error
	AddOrUpdate(cart *models.Cart) error
	RemoveItem(userID, productID uuid.UUID) error
	GetByUserID(userID uuid.UUID) ([]models.Cart, error)
	UpdateQuantity(userID, productID uuid.UUID, quantity int) error
	ToggleIsChecked(userID, productID uuid.UUID) error
}

type cartRepository struct{ db *gorm.DB }

func NewCartRepository(db *gorm.DB) CartRepository {
	return &cartRepository{db}
}

func (r *cartRepository) GetByUserID(userID uuid.UUID) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Product.ProductGallery").
		Where("user_id = ?", userID).
		Find(&carts).Error
	return carts, err
}

func (r *cartRepository) AddOrUpdate(cart *models.Cart) error {
	var existing models.Cart
	err := r.db.Where("user_id = ? AND product_id = ?", cart.UserID, cart.ProductID).
		First(&existing).Error
	if err == nil {
		existing.Quantity += cart.Quantity
		return r.db.Save(&existing).Error
	}
	return r.db.Create(cart).Error
}

func (r *cartRepository) UpdateQuantity(userID, productID uuid.UUID, quantity int) error {
	return r.db.Model(&models.Cart{}).
		Where("user_id = ? AND product_id = ?", userID, productID).
		Update("quantity", quantity).Error
}

func (r *cartRepository) RemoveItem(userID, productID uuid.UUID) error {
	return r.db.Where("user_id = ? AND product_id = ?", userID, productID).
		Delete(&models.Cart{}).Error
}

func (r *cartRepository) Clear(userID uuid.UUID) error {
	return r.db.Where("user_id = ?", userID).
		Delete(&models.Cart{}).Error
}

func (r *cartRepository) ToggleIsChecked(userID, productID uuid.UUID) error {
	var cart models.Cart
	if err := r.db.Where("user_id = ? AND product_id = ?", userID, productID).First(&cart).Error; err != nil {
		return err
	}

	return r.db.Model(&models.Cart{}).
		Where("user_id = ? AND product_id = ?", userID, productID).
		Update("is_checked", !cart.IsChecked).Error
}
