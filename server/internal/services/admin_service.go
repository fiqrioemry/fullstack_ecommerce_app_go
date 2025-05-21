// internal/services/user_service.go
package services

import (
	"server/internal/dto"
	"server/internal/repositories"
	"time"
)

type AdminService interface {
	GetAllCustomer(params dto.CustomerQueryParam) ([]dto.CustomerListResponse, *dto.PaginationResponse, error)
	GetCustomerDetail(id string) (*dto.CustomerDetailResponse, error)
	GetDashboardStats(gender string) (*dto.DashboardStatsResponse, error)
	GetRevenueStats(rangeType string) ([]dto.RevenueStat, float64, error)
}
type adminService struct {
	repo repositories.AdminRepository
}

func NewAdminService(repo repositories.AdminRepository) AdminService {
	return &adminService{repo}
}

func (s *adminService) GetAllCustomer(params dto.CustomerQueryParam) ([]dto.CustomerListResponse, *dto.PaginationResponse, error) {
	customers, total, err := s.repo.FindAllCustomers(params)
	if err != nil {
		return nil, nil, err
	}

	var result []dto.CustomerListResponse
	for _, c := range customers {
		result = append(result, dto.CustomerListResponse{
			ID:        c.ID.String(),
			Email:     c.Email,
			Fullname:  c.Profile.Fullname,
			Avatar:    c.Profile.Avatar,
			CreatedAt: c.CreatedAt.Format(time.RFC3339),
		})
	}
	totalPages := int((total + int64(params.Limit) - 1) / int64(params.Limit))
	return result, &dto.PaginationResponse{
		Page:       params.Page,
		Limit:      params.Limit,
		TotalRows:  int(total),
		TotalPages: totalPages,
	}, nil
}

func (s *adminService) GetCustomerDetail(id string) (*dto.CustomerDetailResponse, error) {
	u, err := s.repo.FindCustomerByID(id)
	if err != nil {
		return nil, err
	}

	res := &dto.CustomerDetailResponse{
		ID:        u.ID.String(),
		Email:     u.Email,
		Fullname:  u.Profile.Fullname,
		Phone:     u.Profile.Phone,
		Avatar:    u.Profile.Avatar,
		Gender:    u.Profile.Gender,
		Address:   u.Addresses[0].Address,
		CreatedAt: u.CreatedAt.Format(time.RFC3339),
		UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
	}
	if u.Profile.Birthday != nil {
		res.Birthday = u.Profile.Birthday.Format("2006-01-02")
	}
	return res, nil
}

// Implementasi fungsi
func (s *adminService) GetDashboardStats(gender string) (*dto.DashboardStatsResponse, error) {
	totalCustomer, err := s.repo.CountCustomerByGender(gender)
	if err != nil {
		return nil, err
	}
	totalProduct, err := s.repo.CountProducts()
	if err != nil {
		return nil, err
	}
	totalOrder, err := s.repo.CountOrders()
	if err != nil {
		return nil, err
	}
	totalRevenue, err := s.repo.SumRevenue()
	if err != nil {
		return nil, err
	}

	return &dto.DashboardStatsResponse{
		TotalCustomers: totalCustomer,
		TotalProducts:  totalProduct,
		TotalOrders:    totalOrder,
		TotalRevenue:   totalRevenue,
	}, nil
}

func (s *adminService) GetRevenueStats(rangeType string) ([]dto.RevenueStat, float64, error) {
	return s.repo.GetRevenueStatsByRange(rangeType)
}
