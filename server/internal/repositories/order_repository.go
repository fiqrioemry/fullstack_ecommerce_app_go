package repositories

import (
	"server/internal/dto"
	"server/internal/models"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	UpdateOrder(order *models.Order) error
	CreateOrderItems(items []models.OrderItem) error
	GetMainAddress(userID uuid.UUID) (*models.Address, error)
	ClearUserCart(userID uuid.UUID) error
	GetUserCart(userID uuid.UUID) ([]models.Cart, error)
	GetOrdersByUserID(userID string, param dto.OrderQueryParam) ([]models.Order, int64, error)
	GetOrderDetail(orderID string) (*models.Order, error)
	GetAllOrders(param dto.OrderQueryParam) ([]models.Order, int64, error)
	MarkOrderDelivered(orderID uuid.UUID) error
	WithTx(fn func(tx *gorm.DB) error) error
	CreateShipment(shipment *models.Shipment) error
	GetShipmentByOrderID(orderID uuid.UUID) (*models.Shipment, error)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &orderRepository{db}
}

func (r *orderRepository) CreateOrder(order *models.Order) error {
	return r.db.Create(order).Error
}

func (r *orderRepository) CreateOrderItems(items []models.OrderItem) error {
	return r.db.Create(&items).Error
}

func (r *orderRepository) GetMainAddress(userID uuid.UUID) (*models.Address, error) {
	var addr models.Address
	err := r.db.Where("user_id = ? AND is_main = ?", userID, true).First(&addr).Error
	return &addr, err
}

func (r *orderRepository) GetUserCart(userID uuid.UUID) ([]models.Cart, error) {
	var carts []models.Cart
	err := r.db.Preload("Product.ProductGallery").
		Where("user_id = ? AND is_checked = ?", userID, true).
		Find(&carts).Error
	return carts, err
}

func (r *orderRepository) ClearUserCart(userID uuid.UUID) error {
	return r.db.Where("user_id = ?", userID).Delete(&models.Cart{}).Error
}

func (r *orderRepository) GetAllOrders(param dto.OrderQueryParam) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	query := r.db.Model(&models.Order{}).Preload("Items")

	if param.Status != "" && param.Status != "all" {
		query = query.Where("status = ?", param.Status)
	}

	if param.Search != "" {
		search := "%" + param.Search + "%"
		query = query.Joins("JOIN order_items ON order_items.order_id = orders.id").
			Where("order_items.product_name LIKE ?", search)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	sort := "created_at desc"
	switch param.Sort {
	case "created_at asc":
		sort = "created_at asc"
	case "created_at desc":
		sort = "created_at desc"
	case "product_name asc":
		sort = "order_items.product_name asc"
	case "product_name desc":
		sort = "order_items.product_name desc"
	}

	offset := (param.Page - 1) * param.Limit
	err := query.
		Order(sort).
		Offset(offset).
		Limit(param.Limit).
		Find(&orders).Error

	return orders, total, err
}

func (r *orderRepository) GetOrdersByUserID(userID string, param dto.OrderQueryParam) ([]models.Order, int64, error) {
	var orders []models.Order
	var total int64

	query := r.db.Model(&models.Order{}).
		Preload("Items").
		Where("user_id = ?", userID)

	if param.Status != "" && param.Status != "all" {
		query = query.Where("status = ?", param.Status)
	}

	if param.Search != "" {
		search := "%" + param.Search + "%"
		query = query.Joins("JOIN order_items ON order_items.order_id = orders.id").
			Where("order_items.product_name LIKE ?", search)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	sort := "created_at desc"
	switch param.Sort {
	case "created_at asc":
		sort = "created_at asc"
	case "created_at desc":
		sort = "created_at desc"
	case "product_name asc":
		sort = "order_items.product_name asc"
	case "product_name desc":
		sort = "order_items.product_name desc"
	}

	offset := (param.Page - 1) * param.Limit
	err := query.
		Order(sort).
		Offset(offset).
		Limit(param.Limit).
		Find(&orders).Error

	return orders, total, err
}

func (r *orderRepository) GetOrderDetail(orderID string) (*models.Order, error) {
	var order models.Order
	err := r.db.Preload("Shipment").Preload("Items").Preload("Address").First(&order, "id = ?", orderID).Error
	return &order, err
}

func (r *orderRepository) CreateShipment(shipment *models.Shipment) error {
	return r.db.Create(shipment).Error
}

func (r *orderRepository) GetShipmentByOrderID(orderID uuid.UUID) (*models.Shipment, error) {
	var shipment models.Shipment
	if err := r.db.First(&shipment, "order_id = ?", orderID).Error; err != nil {
		return nil, err
	}
	return &shipment, nil
}
func (r *orderRepository) WithTx(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r *orderRepository) MarkOrderDelivered(orderID uuid.UUID) error {
	now := time.Now()
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Shipment{}).
			Where("order_id = ?", orderID).
			Updates(map[string]interface{}{
				"status":       "delivered",
				"delivered_at": now,
			}).Error; err != nil {
			return err
		}

		if err := tx.Model(&models.Order{}).
			Where("id = ?", orderID).
			Update("status", "delivered").Error; err != nil {
			return err
		}

		return nil
	})
}

func (r *orderRepository) UpdateOrder(order *models.Order) error {
	return r.db.Model(&models.Order{}).
		Where("id = ?", order.ID).
		Updates(order).Error
}
