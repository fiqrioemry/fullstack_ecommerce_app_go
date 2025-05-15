package repositories

import (
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BannerRepository interface {
	GetAll() ([]models.Banner, error)
	Create(banner *models.Banner) error
	GetByPosition(position string) ([]models.Banner, error)
	Delete(id uuid.UUID) error
	FindByID(id uuid.UUID) (*models.Banner, error)
	Update(b *models.Banner) error
}

type bannerRepo struct {
	db *gorm.DB
}

func NewBannerRepository(db *gorm.DB) BannerRepository {
	return &bannerRepo{db: db}
}

func (r *bannerRepo) Create(banner *models.Banner) error {
	return r.db.Create(banner).Error
}

func (r *bannerRepo) GetAll() ([]models.Banner, error) {
	var banners []models.Banner
	err := r.db.Find(&banners).Error
	return banners, err
}

func (r *bannerRepo) GetByPosition(position string) ([]models.Banner, error) {
	var banners []models.Banner
	err := r.db.Where("position = ?", position).Find(&banners).Error
	return banners, err
}

func (r *bannerRepo) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Banner{}, "id = ?", id).Error
}

func (r *bannerRepo) FindByID(id uuid.UUID) (*models.Banner, error) {
	var b models.Banner
	err := r.db.First(&b, "id = ?", id).Error
	return &b, err
}

func (r *bannerRepo) Update(b *models.Banner) error {
	return r.db.Save(b).Error
}
