package repositories

import (
	"fmt"
	"server/internal/dto"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductRepository interface {
	DeleteProduct(id uuid.UUID) error
	UpdateProduct(product *models.Product) error
	CreateProduct(product *models.Product) error
	GetProductByID(id uuid.UUID) (*models.Product, error)
	CreateProductGallery(image *models.ProductGallery) error
	DeleteProductGalleryByProductID(productID uuid.UUID) error
	GetProductBySlug(slug string) (*models.Product, error)
	DecreaseProductStock(productID uuid.UUID, qty int) error
	RestoreStockOnPaymentFailure(order *models.Order) error
	SearchProducts(param dto.GetAllProductsRequest) ([]models.Product, int64, error)
}

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

func (r *productRepository) CreateProduct(product *models.Product) error {
	return r.db.Create(product).Error
}

func (r *productRepository) CreateProductGallery(image *models.ProductGallery) error {
	return r.db.Create(image).Error
}

func (r *productRepository) UpdateProduct(product *models.Product) error {
	return r.db.Save(product).Error
}

func (r *productRepository) GetProductByID(id uuid.UUID) (*models.Product, error) {
	var product models.Product
	if err := r.db.Preload("ProductGallery").Preload("Category").First(&product, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) DeleteProductGalleryByProductID(productID uuid.UUID) error {
	return r.db.Where("product_id = ?", productID).Delete(&models.ProductGallery{}).Error
}

func (r *productRepository) DeleteProduct(id uuid.UUID) error {
	return r.db.Delete(&models.Product{}, "id = ?", id).Error
}

func (r *productRepository) GetProductBySlug(slug string) (*models.Product, error) {
	var product models.Product
	err := r.db.Preload("ProductGallery").Preload("Category").
		Where("slug = ?", slug).First(&product).Error
	return &product, err
}

func (r *productRepository) SearchProducts(param dto.GetAllProductsRequest) ([]models.Product, int64, error) {
	var products []models.Product
	var total int64

	db := r.db.Model(&models.Product{}).
		Preload("ProductGallery").
		Preload("Category").
		Where("products.is_active = ?", true)

	// Keyword Search
	if param.Search != "" {
		likeQuery := "%" + param.Search + "%"
		db = db.Where("products.name LIKE ? OR products.description LIKE ?", likeQuery, likeQuery)
	}

	// Filter by Category Slug
	if param.Category != "" {
		db = db.Joins("JOIN categories ON categories.id = products.category_id").
			Where("categories.slug = ?", param.Category)
	}
	if param.Status != "" && param.Status != "all" {
		switch param.Status {
		case "active":
			db = db.Where("is_active = ?", true)
		case "inactive":
			db = db.Where("is_active = ?", false)
		case "featured":
			db = db.Where("is_featured = ?", true)
		case "unfeatured":
			db = db.Where("is_featured = ?", false)
		}
	}

	// Price Range Filter
	if param.MinPrice > 0 {
		db = db.Where("products.price >= ?", param.MinPrice)
	}
	if param.MaxPrice > 0 {
		db = db.Where("products.price <= ?", param.MaxPrice)
	}

	// Minimum Rating Filter
	if param.Rating > 0 {
		db = db.Where("products.average_rating >= ?", param.Rating)
	}

	// Sorting
	sort := "products.name asc"
	switch param.Sort {
	case "price_asc":
		sort = "products.price asc"
	case "price_desc":
		sort = "products.price desc"
	case "created_at_asc":
		sort = "products.created_at asc"
	case "created_at_desc":
		sort = "products.created_at desc"
	case "rating_asc":
		sort = "products.average_rating asc"
	case "rating_desc":
		sort = "products.average_rating desc"
	case "name_desc":
		sort = "products.name desc"
	}
	db = db.Order(sort)

	// Pagination
	page := param.Page
	if page <= 0 {
		page = 1
	}
	limit := param.Limit
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	// Get total count
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Get paginated result
	if err := db.Limit(limit).Offset(offset).Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, total, nil
}

func (r *productRepository) DecreaseProductStock(productID uuid.UUID, qty int) error {
	return r.db.Model(&models.Product{}).
		Where("id = ? AND stock >= ?", productID, qty).
		Update("stock", gorm.Expr("stock - ?", qty)).Error
}

func (r *productRepository) RestoreStockOnPaymentFailure(order *models.Order) error {
	for _, item := range order.Items {
		err := r.db.Model(&models.Product{}).
			Where("id = ?", item.ProductID).
			Update("stock", gorm.Expr("stock + ?", item.Quantity)).Error
		if err != nil {
			return fmt.Errorf("failed to restore stock for product ID %s: %w", item.ProductID, err)
		}
	}
	return nil
}
