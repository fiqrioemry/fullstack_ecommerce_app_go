package repositories

import (
	"server/internal/dto"
	"server/internal/models"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	DeleteCategory(CategoryID string) error
	GetAllCategories(param dto.CategoryQueryParam) ([]models.Category, int64, error)
	CreateCategory(category *models.Category) error
	UpdateCategory(category *models.Category) error
	GetCategoryByID(CategoryID string) (*models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (r *categoryRepository) GetAllCategories(param dto.CategoryQueryParam) ([]models.Category, int64, error) {
	var categories []models.Category
	var total int64

	db := r.db.Model(&models.Category{})

	if param.Search != "" {
		search := "%" + param.Search + "%"
		db = db.Where("name LIKE ?", search)
	}

	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	sort := "name asc"
	switch param.Sort {
	case "created_at_asc":
		sort = "created_at asc"
	case "created_at_desc":
		sort = "created_at desc"
	}

	db = db.Order(sort)

	offset := (param.Page - 1) * param.Limit
	if err := db.Offset(offset).Limit(param.Limit).Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

func (r *categoryRepository) CreateCategory(category *models.Category) error {
	return r.db.Create(category).Error
}

func (r *categoryRepository) UpdateCategory(category *models.Category) error {
	return r.db.Save(category).Error
}

func (r *categoryRepository) DeleteCategory(CategoryID string) error {
	return r.db.Delete(&models.Category{}, "id = ?", CategoryID).Error
}

func (r *categoryRepository) GetCategoryByID(CategoryID string) (*models.Category, error) {
	var category models.Category
	if err := r.db.First(&category, "id = ?", CategoryID).Error; err != nil {
		return nil, err
	}
	return &category, nil
}
