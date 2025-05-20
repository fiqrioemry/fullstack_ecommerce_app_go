package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"

	"github.com/google/uuid"
)

type BannerService interface {
	GetAll() ([]dto.BannerResponse, error)
	Create(req dto.BannerRequest) error
	Get(position string) ([]dto.BannerResponse, error)
	Delete(id uuid.UUID) error
	Update(id uuid.UUID, req dto.BannerRequest) error
}

type bannerService struct {
	bannerRepo repositories.BannerRepository
}

func NewBannerService(bannerRepo repositories.BannerRepository) BannerService {
	return &bannerService{bannerRepo}
}

func (s *bannerService) Create(req dto.BannerRequest) error {
	banner := &models.Banner{
		Position: req.Position,
		Image:    req.ImageURL,
	}
	return s.bannerRepo.Create(banner)
}

func (s *bannerService) GetAll() ([]dto.BannerResponse, error) {
	banners, err := s.bannerRepo.GetAll()
	if err != nil {
		return nil, err
	}
	var results []dto.BannerResponse
	for _, b := range banners {
		results = append(results, dto.BannerResponse{
			ID:       b.ID.String(),
			Position: b.Position,
			Image:    b.Image,
		})
	}
	return results, nil
}

func (s *bannerService) Get(position string) ([]dto.BannerResponse, error) {
	banners, err := s.bannerRepo.GetByPosition(position)
	if err != nil {
		return nil, err
	}
	var results []dto.BannerResponse
	for _, b := range banners {
		results = append(results, dto.BannerResponse{
			ID:       b.ID.String(),
			Position: b.Position,
			Image:    b.Image,
		})
	}
	return results, nil
}

func (s *bannerService) Delete(id uuid.UUID) error {
	b, err := s.bannerRepo.FindByID(id)
	if err != nil {
		return err
	}
	_ = utils.DeleteFromCloudinary(b.Image)
	return s.bannerRepo.Delete(id)
}

func (s *bannerService) Update(id uuid.UUID, req dto.BannerRequest) error {
	banner, err := s.bannerRepo.FindByID(id)
	if err != nil {
		return err
	}

	banner.Image = req.ImageURL
	banner.Position = req.Position
	return s.bannerRepo.Update(banner)
}
