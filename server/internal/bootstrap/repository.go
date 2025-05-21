package bootstrap

import (
	"server/internal/repositories"

	"gorm.io/gorm"
)

type RepositoryContainer struct {
	AdminRepository        repositories.AdminRepository
	AuthRepository         repositories.AuthRepository
	VoucherRepository      repositories.VoucherRepository
	ProductRepository      repositories.ProductRepository
	PaymentRepository      repositories.PaymentRepository
	ProfileRepository      repositories.ProfileRepository
	CartRepository         repositories.CartRepository
	OrderRepository        repositories.OrderRepository
	LocationRepository     repositories.LocationRepository
	AddressRepository      repositories.AddressRepository
	CategoryRepository     repositories.CategoryRepository
	NotificationRepository repositories.NotificationRepository
	BannerRepository       repositories.BannerRepository
	ReviewRepository       repositories.ReviewRepository
}

func InitRepositories(db *gorm.DB) *RepositoryContainer {
	return &RepositoryContainer{
		AdminRepository:        repositories.NewAdminRepository(db),
		AuthRepository:         repositories.NewAuthRepository(db),
		ProductRepository:      repositories.NewProductRepository(db),
		PaymentRepository:      repositories.NewPaymentRepository(db),
		VoucherRepository:      repositories.NewVoucherRepository(db),
		ProfileRepository:      repositories.NewProfileRepository(db),
		CartRepository:         repositories.NewCartRepository(db),
		OrderRepository:        repositories.NewOrderRepository(db),
		LocationRepository:     repositories.NewLocationRepository(db),
		AddressRepository:      repositories.NewAddressRepository(db),
		CategoryRepository:     repositories.NewCategoryRepository(db),
		NotificationRepository: repositories.NewNotificationRepository(db),
		BannerRepository:       repositories.NewBannerRepository(db),
		ReviewRepository:       repositories.NewReviewRepository(db),
	}
}
