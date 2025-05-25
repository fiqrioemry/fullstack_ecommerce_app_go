package repositories

import (
	"server/internal/dto"
	"server/internal/models"

	"gorm.io/gorm"
)

type AddressRepository interface {
	CreateAddress(address *models.Address) error
	UpdateAddress(address *models.Address) error
	DeleteAddress(addressID string, userID string) error
	SetMainAddress(addressID string, userID string) error
	UnsetAllMain(userID string) error
	GetMainAddress(userID string) (*models.Address, error)
	GetAddressByID(addressID string) (*models.Address, error)
	GetAddressesByUserID(userID string, param dto.AddressQueryParam) ([]models.Address, int64, error)
}

type addressRepo struct {
	db *gorm.DB
}

func NewAddressRepository(db *gorm.DB) AddressRepository {
	return &addressRepo{db}
}

func (r *addressRepo) GetAddressesByUserID(userID string, params dto.AddressQueryParam) ([]models.Address, int64, error) {
	var addresses []models.Address
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
	query := r.db.Model(&models.Address{}).
		Where("user_id = ?", userID)

	// Pencarian keyword
	if params.Q != "" {
		q := "%" + params.Q + "%"
		query = query.Where("name LIKE ? OR address LIKE ?", q, q)
	}

	// Sorting
	sort := "is_main desc" // default prioritaskan alamat utama
	switch params.Sort {
	case "name_asc":
		sort = "name asc"
	case "name_desc":
		sort = "name desc"
	case "created_at_asc":
		sort = "created_at asc"
	case "created_at_desc":
		sort = "created_at desc"
	}
	query = query.Order(sort)

	// Hitung total
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// Ambil hasil paginasi
	if err := query.Offset(offset).Limit(limit).Find(&addresses).Error; err != nil {
		return nil, 0, err
	}

	return addresses, total, nil
}

func (r *addressRepo) CreateAddress(address *models.Address) error {
	return r.db.Create(address).Error
}

func (r *addressRepo) UpdateAddress(address *models.Address) error {
	return r.db.Save(address).Error
}

func (r *addressRepo) DeleteAddress(addressID string, userID string) error {
	return r.db.Where("id = ? AND user_id = ?", addressID, userID).Delete(&models.Address{}).Error
}

func (r *addressRepo) SetMainAddress(addressID string, userID string) error {
	return r.db.Model(&models.Address{}).Where("id = ? AND user_id = ?", addressID, userID).Update("is_main", true).Error
}

func (r *addressRepo) UnsetAllMain(userID string) error {
	return r.db.Model(&models.Address{}).Where("user_id = ?", userID).Update("is_main", false).Error
}

func (r *addressRepo) GetAddressByID(addressID string) (*models.Address, error) {
	var addr models.Address
	err := r.db.Where("id = ? ", addressID).Find(&addr).Error
	return &addr, err
}

func (r *addressRepo) GetMainAddress(userID string) (*models.Address, error) {
	var addr models.Address
	err := r.db.Where("user_id = ? AND is_main =?", userID, true).Find(&addr).Error
	return &addr, err
}
