package bootstrap

import (
	"server/internal/handlers"
)

type HandlerContainer struct {
	AuthHandler         *handlers.AuthHandler
	VoucherHandler      *handlers.VoucherHandler
	ProductHandler      *handlers.ProductHandler
	PaymentHandler      *handlers.PaymentHandler
	ProfileHandler      *handlers.ProfileHandler
	CartHandler         *handlers.CartHandler
	OrderHandler        *handlers.OrderHandler
	AddressHandler      *handlers.AddressHandler
	LocationHandler     *handlers.LocationHandler
	CategoryHandler     *handlers.CategoryHandler
	NotificationHandler *handlers.NotificationHandler
	BannerHandler       *handlers.BannerHandler
	ReviewHandler       *handlers.ReviewHandler
}

func InitHandlers(svc *ServiceContainer) *HandlerContainer {
	return &HandlerContainer{
		AuthHandler:         handlers.NewAuthHandler(svc.AuthService),
		ProductHandler:      handlers.NewProductHandler(svc.ProductService),
		VoucherHandler:      handlers.NewVoucherHandler(svc.VoucherService),
		PaymentHandler:      handlers.NewPaymentHandler(svc.PaymentService),
		ProfileHandler:      handlers.NewProfileHandler(svc.ProfileService),
		CartHandler:         handlers.NewCartHandler(svc.CartService),
		OrderHandler:        handlers.NewOrderHandler(svc.OrderService),
		LocationHandler:     handlers.NewLocationHandler(svc.LocationService),
		AddressHandler:      handlers.NewAddressHandler(svc.AddressService),
		CategoryHandler:     handlers.NewCategoryHandler(svc.CategoryService),
		NotificationHandler: handlers.NewNotificationHandler(svc.NotificationService),
		BannerHandler:       handlers.NewBannerHandler(svc.BannerService),
		ReviewHandler:       handlers.NewReviewHandler(svc.ReviewService),
	}
}
