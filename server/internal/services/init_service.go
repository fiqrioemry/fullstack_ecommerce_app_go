package services

import (
	"server/internal/repositories"

)

type Services struct {
	AdminService        AdminService
	AuthService         AuthService
	ProductService      ProductService
	VoucherService      VoucherService
	BannerService       BannerService
	ProfileService      ProfileService
	PaymentService      PaymentService
	CartService         CartService
	OrderService        OrderService
	AddressService      AddressService
	LocationService     LocationService
	CategoryService     CategoryService
	NotificationService NotificationService
	ReviewService       ReviewService
}

func InitServices(r *repositories.Repositories) *Services {
	voucherSvc := NewVoucherService(r.VoucherRepository)
	notificationSvc := NewNotificationService(r.NotificationRepository)
	return &Services{
		VoucherService:      voucherSvc,
		AdminService:        NewAdminService(r.AdminRepository),
		BannerService:       NewBannerService(r.BannerRepository),
		ProductService:      NewProductService(r.ProductRepository),
		ProfileService:      NewProfileService(r.ProfileRepository),
		LocationService:     NewLocationService(r.LocationRepository),
		CategoryService:     NewCategoryService(r.CategoryRepository),
		NotificationService: NewNotificationService(r.NotificationRepository),
		CartService:         NewCartService(r.CartRepository, r.ProductRepository),
		AuthService:         NewAuthService(r.AuthRepository, r.NotificationRepository),
		AddressService:      NewAddressService(r.AddressRepository, r.LocationRepository),
		PaymentService:      NewPaymentService(r.PaymentRepository, r.AuthRepository, r.ProductRepository, voucherSvc, r.OrderRepository, notificationSvc),
		OrderService:        NewOrderService(r.OrderRepository, r.PaymentRepository, r.AuthRepository, r.ProductRepository, voucherSvc, notificationSvc),
		ReviewService:       NewReviewService(r.ReviewRepository, r.OrderRepository),
	}
}
