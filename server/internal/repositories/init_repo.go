package repositories

import (
	"gorm.io/gorm"
)

type Repositories struct {
	AdminRepository        AdminRepository
	AuthRepository         AuthRepository
	VoucherRepository      VoucherRepository
	ProductRepository      ProductRepository
	PaymentRepository      PaymentRepository
	ProfileRepository      ProfileRepository
	CartRepository         CartRepository
	OrderRepository        OrderRepository
	LocationRepository     LocationRepository
	AddressRepository      AddressRepository
	CategoryRepository     CategoryRepository
	NotificationRepository NotificationRepository
	BannerRepository       BannerRepository
	ReviewRepository       ReviewRepository
}

func InitRepositories(db *gorm.DB) *Repositories {
	return &Repositories{
		AdminRepository:        NewAdminRepository(db),
		AuthRepository:         NewAuthRepository(db),
		ProductRepository:      NewProductRepository(db),
		PaymentRepository:      NewPaymentRepository(db),
		VoucherRepository:      NewVoucherRepository(db),
		ProfileRepository:      NewProfileRepository(db),
		CartRepository:         NewCartRepository(db),
		OrderRepository:        NewOrderRepository(db),
		LocationRepository:     NewLocationRepository(db),
		AddressRepository:      NewAddressRepository(db),
		CategoryRepository:     NewCategoryRepository(db),
		NotificationRepository: NewNotificationRepository(db),
		BannerRepository:       NewBannerRepository(db),
		ReviewRepository:       NewReviewRepository(db),
	}
}
