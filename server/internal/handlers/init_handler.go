package handlers

import (
	"server/internal/services"
)

type Handlers struct {
	AdminHandler        *AdminHandler
	AuthHandler         *AuthHandler
	VoucherHandler      *VoucherHandler
	ProductHandler      *ProductHandler
	PaymentHandler      *PaymentHandler
	ProfileHandler      *ProfileHandler
	CartHandler         *CartHandler
	OrderHandler        *OrderHandler
	AddressHandler      *AddressHandler
	LocationHandler     *LocationHandler
	CategoryHandler     *CategoryHandler
	NotificationHandler *NotificationHandler
	BannerHandler       *BannerHandler
	ReviewHandler       *ReviewHandler
}

func InitHandlers(s *services.Services) *Handlers{
	return &Handlers{
		AdminHandler:        NewAdminHandler(s.AdminService),
		AuthHandler:         NewAuthHandler(s.AuthService),
		ProductHandler:      NewProductHandler(s.ProductService),
		VoucherHandler:      NewVoucherHandler(s.VoucherService),
		PaymentHandler:      NewPaymentHandler(s.PaymentService),
		ProfileHandler:      NewProfileHandler(s.ProfileService),
		CartHandler:         NewCartHandler(s.CartService),
		OrderHandler:        NewOrderHandler(s.OrderService),
		LocationHandler:     NewLocationHandler(s.LocationService),
		AddressHandler:      NewAddressHandler(s.AddressService),
		CategoryHandler:     NewCategoryHandler(s.CategoryService),
		NotificationHandler: NewNotificationHandler(s.NotificationService),
		BannerHandler:       NewBannerHandler(s.BannerService),
		ReviewHandler:       NewReviewHandler(s.ReviewService),
	}
}
