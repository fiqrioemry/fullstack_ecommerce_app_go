package services

import (
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"
)

type CategoryService interface {
	GetAllCategories(param dto.CategoryQueryParam) ([]dto.CategoryListResponse, *dto.PaginationResponse, error)
	DeleteCategory(categoryID string) error
	CreateCategory(req dto.CreateCategoryRequest) error
	GetCategoryByID(categoryID string) (*dto.CategoryResponse, error)
	UpdateCategory(categoryID string, req dto.UpdateCategoryRequest) error
}

type categoryService struct {
	repo repositories.CategoryRepository
}

func NewCategoryService(repo repositories.CategoryRepository) CategoryService {
	return &categoryService{repo}
}
func (s *categoryService) GetAllCategories(param dto.CategoryQueryParam) ([]dto.CategoryListResponse, *dto.PaginationResponse, error) {
	categories, total, err := s.repo.GetAllCategories(param)
	if err != nil {
		return nil, nil, err
	}

	var result []dto.CategoryListResponse
	for _, c := range categories {
		result = append(result, dto.CategoryListResponse{
			ID:    c.ID.String(),
			Name:  c.Name,
			Slug:  c.Slug,
			Image: c.Image,
		})
	}

	totalPages := int((total + int64(param.Limit) - 1) / int64(param.Limit))
	pagination := &dto.PaginationResponse{
		Page:       param.Page,
		Limit:      param.Limit,
		TotalRows:  int(total),
		TotalPages: totalPages,
	}

	return result, pagination, nil
}

func (s *categoryService) CreateCategory(req dto.CreateCategoryRequest) error {
	category := models.Category{
		Name:  req.Name,
		Slug:  utils.GenerateSlug(req.Name),
		Image: req.ImageURL,
	}
	return s.repo.CreateCategory(&category)
}

func (s *categoryService) UpdateCategory(categoryID string, req dto.UpdateCategoryRequest) error {
	category, err := s.repo.GetCategoryByID(categoryID)
	if err != nil {
		return err
	}

	if req.ImageURL != "" {
		category.Name = req.Name
		category.Slug = utils.GenerateSlug(req.Name)
	}

	if req.ImageURL != "" {
		category.Image = req.ImageURL
	}

	return s.repo.UpdateCategory(category)
}

func (s *categoryService) DeleteCategory(categoryID string) error {
	category, err := s.repo.GetCategoryByID(categoryID)
	if err != nil {
		return err
	}

	utils.CleanupImageOnError(category.Image)

	return s.repo.DeleteCategory(categoryID)
}

func (s *categoryService) GetCategoryByID(categoryID string) (*dto.CategoryResponse, error) {
	category, err := s.repo.GetCategoryByID(categoryID)
	if err != nil {
		return nil, err
	}

	return &dto.CategoryResponse{
		ID:    categoryID,
		Name:  category.Name,
		Slug:  category.Slug,
		Image: category.Image,
	}, nil
}
