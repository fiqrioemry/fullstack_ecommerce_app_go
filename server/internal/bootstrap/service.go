package bootstrap

import (
	"server/internal/services"
)

type ServiceContainer struct {
	AuthService         services.AuthService
	ProductService      services.ProductService
	VoucherService      services.VoucherService
	BannerService       services.BannerService
	ProfileService      services.ProfileService
	PaymentService      services.PaymentService
	CartService         services.CartService
	OrderService        services.OrderService
	AddressService      services.AddressService
	LocationService     services.LocationService
	CategoryService     services.CategoryService
	NotificationService services.NotificationService
	ReviewService       services.ReviewService
}

func InitServices(repo *RepositoryContainer) *ServiceContainer {
	voucherSvc := services.NewVoucherService(repo.VoucherRepository)
	return &ServiceContainer{
		VoucherService:      voucherSvc,
		BannerService:       services.NewBannerService(repo.BannerRepository),
		ProductService:      services.NewProductService(repo.ProductRepository),
		ProfileService:      services.NewProfileService(repo.ProfileRepository),
		LocationService:     services.NewLocationService(repo.LocationRepository),
		CategoryService:     services.NewCategoryService(repo.CategoryRepository),
		NotificationService: services.NewNotificationService(repo.NotificationRepository),
		CartService:         services.NewCartService(repo.CartRepository, repo.ProductRepository),
		AuthService:         services.NewAuthService(repo.AuthRepository, repo.NotificationRepository),
		AddressService:      services.NewAddressService(repo.AddressRepository, repo.LocationRepository),
		PaymentService:      services.NewPaymentService(repo.PaymentRepository, repo.AuthRepository, repo.ProductRepository, voucherSvc, repo.OrderRepository),
		OrderService:        services.NewOrderService(repo.OrderRepository, repo.PaymentRepository, repo.AuthRepository, repo.ProductRepository, voucherSvc),
		ReviewService:       services.NewReviewService(repo.ReviewRepository, repo.OrderRepository),
	}
}
