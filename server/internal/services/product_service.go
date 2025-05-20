package services

import (
	"errors"
	"server/internal/dto"
	"server/internal/models"
	"server/internal/repositories"
	"server/internal/utils"

	"github.com/google/uuid"
)

type ProductService interface {
	DeleteProduct(productID string) error
	CreateProduct(req dto.CreateProductRequest) error
	UpdateProduct(productID string, req dto.UpdateProductRequest) error
	GetProductBySlug(slug string) (*dto.ProductDetailResponse, error)
	SearchProducts(param dto.GetAllProductsRequest) ([]dto.ProductListResponse, *dto.PaginationResponse, error)
}

type productService struct {
	productRepo repositories.ProductRepository
}

func NewProductService(productRepo repositories.ProductRepository) ProductService {
	return &productService{productRepo}
}

func (s *productService) CreateProduct(req dto.CreateProductRequest) error {
	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return errors.New("invalid category ID")
	}

	product := models.Product{
		Name:        req.Name,
		Slug:        utils.GenerateSlug(req.Name),
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
		IsActive:    req.IsActive,
		IsFeatured:  req.IsFeatured,
		CategoryID:  categoryID,
		Discount:    req.Discount,
	}

	if err := s.productRepo.CreateProduct(&product); err != nil {
		return err
	}

	for _, url := range req.ImageURLs {
		img := models.ProductGallery{
			ProductID: product.ID,
			Image:     url,
		}
		if err := s.productRepo.CreateProductGallery(&img); err != nil {
			return err
		}
	}

	return nil
}

func (s *productService) UpdateProduct(productID string, req dto.UpdateProductRequest) error {
	id, err := uuid.Parse(productID)
	if err != nil {
		return errors.New("invalid product ID")
	}

	existingProduct, err := s.productRepo.GetProductByID(id)
	if err != nil {
		return err
	}

	// Handle image update
	if len(req.ImageURLs) > 0 {
		for _, img := range existingProduct.ProductGallery {
			utils.CleanupImageOnError(img.Image)
		}

		if err := s.productRepo.DeleteProductGalleryByProductID(id); err != nil {
			return err
		}

		for _, url := range req.ImageURLs {
			if err := s.productRepo.CreateProductGallery(&models.ProductGallery{
				ProductID: id,
				Image:     url,
			}); err != nil {
				return err
			}
		}
	}

	// Update fields
	categoryID, err := uuid.Parse(req.CategoryID)
	if err != nil {
		return errors.New("invalid category ID")
	}

	existingProduct.Name = req.Name
	existingProduct.Slug = utils.GenerateSlug(req.Name)
	existingProduct.Description = req.Description
	existingProduct.Price = req.Price
	existingProduct.Stock = req.Stock
	existingProduct.Weight = req.Weight
	existingProduct.Length = req.Length
	existingProduct.Width = req.Width
	existingProduct.Height = req.Height
	existingProduct.Discount = req.Discount
	existingProduct.IsActive = req.IsActive
	existingProduct.IsFeatured = req.IsFeatured
	existingProduct.CategoryID = categoryID

	return s.productRepo.UpdateProduct(existingProduct)
}

func (s *productService) DeleteProduct(productID string) error {
	id, err := uuid.Parse(productID)
	if err != nil {
		return errors.New("invalid product ID")
	}

	existing, err := s.productRepo.GetProductByID(id)
	if err != nil {
		return err
	}

	for _, img := range existing.ProductGallery {
		utils.CleanupImageOnError(img.Image)
	}

	return s.productRepo.DeleteProduct(id)
}

func (s *productService) GetProductBySlug(slug string) (*dto.ProductDetailResponse, error) {
	product, err := s.productRepo.GetProductBySlug(slug)
	if err != nil {
		return nil, err
	}
	var images []string
	for _, img := range product.ProductGallery {
		images = append(images, img.Image)
	}
	return &dto.ProductDetailResponse{
		ID:            product.ID.String(),
		Name:          product.Name,
		Slug:          product.Slug,
		Description:   product.Description,
		Price:         product.Price,
		Stock:         product.Stock,
		Discount:      product.Discount,
		Category:      product.Category.Name,
		AverageRating: product.AverageRating,
		Images:        images,
	}, nil
}

func (s *productService) SearchProducts(param dto.GetAllProductsRequest) ([]dto.ProductListResponse, *dto.PaginationResponse, error) {
	products, total, err := s.productRepo.SearchProducts(param)
	if err != nil {
		return nil, nil, err
	}

	var result []dto.ProductListResponse
	for _, p := range products {
		var imageURLs []string
		for _, g := range p.ProductGallery {
			imageURLs = append(imageURLs, g.Image)
		}

		result = append(result, dto.ProductListResponse{
			ID:         p.ID.String(),
			Name:       p.Name,
			Slug:       p.Slug,
			Stock:      p.Stock,
			Price:      p.Price,
			Discount:   p.Discount,
			IsActive:   p.IsActive,
			Category:   p.Category.Name,
			IsFeatured: p.IsFeatured,
			Images:     imageURLs,
		})
	}

	totalPages := int((total + int64(param.Limit) - 1) / int64(param.Limit))
	return result, &dto.PaginationResponse{
		Page:       param.Page,
		Limit:      param.Limit,
		TotalRows:  int(total),
		TotalPages: totalPages,
	}, nil
}
