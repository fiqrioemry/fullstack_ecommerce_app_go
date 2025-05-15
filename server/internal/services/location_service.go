package services

import (
	"server/internal/models"
	"server/internal/repositories"
)

type LocationService interface {
	GetAllProvinces() ([]models.Province, error)
	GetCitiesByProvinceID(provinceID uint) ([]models.City, error)
	SearchCitiesByName(query string) ([]models.City, error)
	GetDistrictsByCityID(cityID uint) ([]models.District, error)
	SearchProvincesByName(query string) ([]models.Province, error)
	GetSubdistrictsByDistrictID(districtID uint) ([]models.Subdistrict, error)
	GetPostalCodesBySubdistrictID(subdistrictID uint) ([]models.PostalCode, error)
}

type locationService struct {
	repo repositories.LocationRepository
}

func NewLocationService(repo repositories.LocationRepository) LocationService {
	return &locationService{repo: repo}
}

func (s *locationService) GetAllProvinces() ([]models.Province, error) {
	return s.repo.GetAllProvinces()
}

func (s *locationService) GetCitiesByProvinceID(provinceID uint) ([]models.City, error) {
	return s.repo.GetCitiesByProvinceID(provinceID)
}

func (s *locationService) SearchCitiesByName(query string) ([]models.City, error) {
	return s.repo.SearchCitiesByName(query)
}

func (s *locationService) GetDistrictsByCityID(cityID uint) ([]models.District, error) {
	return s.repo.GetDistrictsByCityID(cityID)
}

func (s *locationService) GetSubdistrictsByDistrictID(districtID uint) ([]models.Subdistrict, error) {
	return s.repo.GetSubdistrictsByDistrictID(districtID)
}

func (s *locationService) GetPostalCodesBySubdistrictID(subdistrictID uint) ([]models.PostalCode, error) {
	return s.repo.GetPostalCodesBySubdistrictID(subdistrictID)
}

func (s *locationService) SearchProvincesByName(query string) ([]models.Province, error) {
	return s.repo.SearchProvincesByName(query)
}
