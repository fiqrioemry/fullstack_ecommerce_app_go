package repositories

import (
	"server/internal/models"

	"gorm.io/gorm"
)

type LocationRepository interface {
	GetAllProvinces() ([]models.Province, error)
	SearchProvincesByName(query string) ([]models.Province, error)
	GetCitiesByProvinceID(provinceID uint) ([]models.City, error)
	SearchCitiesByName(query string) ([]models.City, error)
	GetDistrictsByCityID(cityID uint) ([]models.District, error)
	GetSubdistrictsByDistrictID(districtID uint) ([]models.Subdistrict, error)
	GetPostalCodesBySubdistrictID(subdistrictID uint) ([]models.PostalCode, error)
}

type locationRepository struct {
	db *gorm.DB
}

func NewLocationRepository(db *gorm.DB) LocationRepository {
	return &locationRepository{db: db}
}

func (r *locationRepository) GetAllProvinces() ([]models.Province, error) {
	var provinces []models.Province
	err := r.db.Order("name ASC").Find(&provinces).Error
	return provinces, err
}

func (r *locationRepository) SearchProvincesByName(query string) ([]models.Province, error) {
	var provinces []models.Province
	err := r.db.Where("name LIKE ?", "%"+query+"%").Order("name ASC").Limit(10).Find(&provinces).Error
	return provinces, err
}

func (r *locationRepository) GetCitiesByProvinceID(provinceID uint) ([]models.City, error) {
	var cities []models.City
	err := r.db.Where("province_id = ?", provinceID).Order("name ASC").Find(&cities).Error
	return cities, err
}

func (r *locationRepository) SearchCitiesByName(query string) ([]models.City, error) {
	var cities []models.City
	err := r.db.Where("name LIKE ?", "%"+query+"%").Order("name ASC").Limit(10).Find(&cities).Error
	return cities, err
}

func (r *locationRepository) GetDistrictsByCityID(cityID uint) ([]models.District, error) {
	var districts []models.District
	err := r.db.Where("city_id = ?", cityID).Order("name ASC").Find(&districts).Error
	return districts, err
}

func (r *locationRepository) GetSubdistrictsByDistrictID(districtID uint) ([]models.Subdistrict, error) {
	var subdistricts []models.Subdistrict
	err := r.db.Where("district_id = ?", districtID).Order("name ASC").Find(&subdistricts).Error
	return subdistricts, err
}

func (r *locationRepository) GetPostalCodesBySubdistrictID(subdistrictID uint) ([]models.PostalCode, error) {
	var postalCodes []models.PostalCode
	err := r.db.Where("subdistrict_id = ?", subdistrictID).Order("postal_code ASC").Find(&postalCodes).Error
	return postalCodes, err
}
