package seeders

import (
	"log"
	"server/internal/models"

	"gorm.io/gorm"
)

func ResetDatabase(db *gorm.DB) {
	log.Println("⚠️ Dropping all tables...")

	err := db.Migrator().DropTable(
		&models.Token{},
		&models.Profile{},
		&models.User{},
		&models.Category{},
		&models.Product{},
		&models.Cart{},
		&models.Banner{},
		&models.Order{},
		&models.OrderItem{},
		&models.Shipment{},
		&models.Address{},
		&models.Province{},
		&models.City{},
		&models.District{},
		&models.Subdistrict{},
		&models.PostalCode{},
		&models.ProductGallery{},
		&models.Payment{},
		&models.Notification{},
		&models.NotificationType{},
		&models.NotificationSetting{},
		&models.Voucher{},
		&models.UsedVoucher{},
		&models.Review{},
	)
	if err != nil {
		log.Fatalf("❌ Failed to drop tables: %v", err)
	}

	log.Println("✅ All tables dropped successfully.")

	log.Println("🚀 Migrating tables...")

	err = db.AutoMigrate(
		&models.Token{},
		&models.Profile{},
		&models.User{},
		&models.Banner{},
		&models.Category{},
		&models.Product{},
		&models.Cart{},
		&models.Order{},
		&models.OrderItem{},
		&models.Shipment{},
		&models.Address{},
		&models.Province{},
		&models.City{},
		&models.District{},
		&models.Subdistrict{},
		&models.PostalCode{},
		&models.ProductGallery{},
		&models.Payment{},
		&models.Notification{},
		&models.NotificationType{},
		&models.NotificationSetting{},
		&models.Voucher{},
		&models.UsedVoucher{},
		&models.Review{},
	)
	if err != nil {
		log.Fatalf("❌ Failed to migrate tables: %v", err)
	}

	log.Println("✅ Migration completed successfully.")

	log.Println("🌱 Seeding dummy data...")
	SeedNotificationTypes(db)
	SeedUsers(db)
	seedProvinces(db)
	seedCities(db)
	seedDistricts(db)
	seedSubdistricts(db)
	seedPostalCodes(db)
	SeedBanner(db)
	SeedCategories(db)
	SeedFashionAndApparel(db)
	SeedFoodBeverage(db)
	SeedGadgetElectronic(db)
	SeedMenShoes(db)
	SeedVouchers(db)
	SeedReviews(db)
	SeedCustomerTransactions(db)
	SeedCustomerNotifications(db)
	log.Println("✅ Seeding completed successfully.")
}
