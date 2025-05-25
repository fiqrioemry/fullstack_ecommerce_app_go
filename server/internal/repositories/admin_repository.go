package repositories

import (
	"server/internal/dto"
	"server/internal/models"

	"gorm.io/gorm"
)

type AdminRepository interface {
	CountProducts() (int64, error)
	CountOrders() (int64, error)
	SumRevenue() (float64, error)
	FindCustomerByID(id string) (*models.User, error)
	CountCustomerByGender(gender string) (int64, error)
	GetRevenueStatsByRange(rangeType string) ([]dto.RevenueStat, float64, error)
	FindAllCustomers(params dto.CustomerQueryParam) ([]models.User, int64, error)
}
type adminRepository struct {
	db *gorm.DB
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db}
}

func (r *adminRepository) FindAllCustomers(params dto.CustomerQueryParam) ([]models.User, int64, error) {
	var users []models.User
	var total int64

	// Validasi pagination
	page := params.Page
	if page <= 0 {
		page = 1
	}
	limit := params.Limit
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	// Base query
	db := r.db.Model(&models.User{}).
		Joins("JOIN profiles ON users.id = profiles.user_id").
		Preload("Profile").
		Where("users.role = ?", "customer")

	// Search query
	if params.Q != "" {
		q := "%" + params.Q + "%"
		db = db.Where("users.email LIKE ? OR profiles.fullname LIKE ?", q, q)
	}

	// Sorting
	sort := "profiles.fullname asc"
	switch params.Sort {
	case "created_at_asc":
		sort = "users.created_at asc"
	case "created_at_desc":
		sort = "users.created_at desc"
	case "name_desc":
		sort = "profiles.fullname desc"
	}
	db = db.Order(sort)

	// Count
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Paginated result
	if err := db.Offset(offset).Limit(limit).Find(&users).Error; err != nil {
		return nil, 0, err
	}

	return users, total, nil
}

func (r *adminRepository) FindCustomerByID(id string) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Profile").
		Preload("Addresses", "is_main = ?", true).
		Preload("Tokens", func(db *gorm.DB) *gorm.DB {
			return db.Order("created_at desc").Limit(1)
		}).
		First(&user, "users.id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *adminRepository) CountCustomerByGender(gender string) (int64, error) {
	var count int64
	db := r.db.Model(&models.User{}).
		Joins("JOIN profiles ON profiles.user_id = users.id").
		Where("users.role = ?", "customer")
	if gender != "" {
		db = db.Where("profiles.gender = ?", gender)
	}
	err := db.Count(&count).Error
	return count, err
}

func (r *adminRepository) CountProducts() (int64, error) {
	var count int64
	err := r.db.Model(&models.Product{}).Count(&count).Error
	return count, err
}

func (r *adminRepository) CountOrders() (int64, error) {
	var count int64
	err := r.db.Model(&models.Order{}).
		Where("status IN ?", []string{"pending", "success"}).
		Count(&count).Error
	return count, err
}

func (r *adminRepository) SumRevenue() (float64, error) {
	var total float64
	err := r.db.Model(&models.Payment{}).
		Where("status = ?", "success").
		Select("SUM(total)").Scan(&total).Error
	return total, err
}

func (r *adminRepository) GetRevenueStatsByRange(rangeType string) ([]dto.RevenueStat, float64, error) {
	var stats []dto.RevenueStat
	var total float64

	query := r.db.Model(&models.Payment{}).Where("status = ?", "success")

	selectClause := ""
	groupClause := ""
	orderClause := ""

	switch rangeType {
	case "daily":
		selectClause = "DATE(paid_at) as date"
		groupClause = "DATE(paid_at)"
		orderClause = "DATE(paid_at)"
	case "monthly":
		selectClause = "DATE_FORMAT(paid_at, '%Y-%m') as date"
		groupClause = "DATE_FORMAT(paid_at, '%Y-%m')"
		orderClause = "DATE_FORMAT(paid_at, '%Y-%m')"
	case "yearly":
		selectClause = "YEAR(paid_at) as date"
		groupClause = "YEAR(paid_at)"
		orderClause = "YEAR(paid_at)"
	default:
		selectClause = "DATE(paid_at) as date"
		groupClause = "DATE(paid_at)"
		orderClause = "DATE(paid_at)"
	}

	err := query.
		Select(selectClause + ", SUM(total) as total").
		Group(groupClause).
		Order(orderClause + " ASC").
		Scan(&stats).Error
	if err != nil {
		return nil, 0, err
	}

	err = query.Select("SUM(total)").Scan(&total).Error
	return stats, total, err
}
