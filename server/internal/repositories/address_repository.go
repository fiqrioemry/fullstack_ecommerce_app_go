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

func (r *addressRepo) GetAddressesByUserID(userID string, param dto.AddressQueryParam) ([]models.Address, int64, error) {
	var addresses []models.Address
	var total int64

	query := r.db.Model(&models.Address{}).Where("user_id = ?", userID)

	// search
	if param.Q != "" {
		like := "%" + param.Q + "%"
		query = query.Where("name LIKE ? OR address LIKE ?", like, like)
	}

	// count total before pagination
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// sorting
	sort := "is_main DESC"
	if param.Sort != "" {
		switch param.Sort {
		case "created_at asc":
			sort += ", created_at ASC"
		case "created_at desc":
			sort += ", created_at DESC"
		case "name asc":
			sort += ", name ASC"
		case "name desc":
			sort += ", name DESC"
		}
	}

	offset := (param.Page - 1) * param.Limit
	err := query.Order(sort).
		Offset(offset).
		Limit(param.Limit).
		Find(&addresses).Error

	return addresses, total, err
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
