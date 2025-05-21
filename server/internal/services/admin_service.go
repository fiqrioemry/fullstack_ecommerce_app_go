// internal/services/user_service.go
package services

import (
	"server/internal/dto"
	"server/internal/repositories"
)

type AdminService interface {
	GetAllCustomer(params dto.CustomerQueryParam) ([]dto.CustomerListResponse, *dto.PaginationResponse, error)
	GetCustomerDetail(id string) (*dto.CustomerDetailResponse, error)
	GetDashboardStats(gender string) (*dto.DashboardStatsResponse, error)
	GetRevenueStats(rangeType string) (*dto.RevenueStatsResponse, error)
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
			CreatedAt: c.CreatedAt.Format("2006-01-02"),
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
		Address:   "",
		CreatedAt: u.CreatedAt.Format("2006-01-02"),
	}

	if len(u.Addresses) > 0 {
		res.Address = u.Addresses[0].Address
	}
	if u.Profile.Birthday != nil {
		res.Birthday = u.Profile.Birthday.Format("2006-01-02")
	}
	if len(u.Tokens) > 0 {
		res.LastLogin = u.Tokens[0].CreatedAt.Format("2006-01-02 15:04:05")
	}

	return res, nil
}

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

func (s *adminService) GetRevenueStats(rangeType string) (*dto.RevenueStatsResponse, error) {
	stats, total, err := s.repo.GetRevenueStatsByRange(rangeType)
	if err != nil {
		return nil, err
	}

	return &dto.RevenueStatsResponse{
		Range:         rangeType,
		TotalRevenue:  total,
		RevenueSeries: stats,
	}, nil
}
