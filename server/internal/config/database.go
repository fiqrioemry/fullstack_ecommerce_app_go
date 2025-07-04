package config

import (
	"fmt"
	"os"
	"time"

	"server/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dbRootURL := os.Getenv("DB_ROOT_URL")
	dbURL := os.Getenv("DB_URL")
	dbName := os.Getenv("DB_NAME")

	// create root connection
	dbRoot, err := gorm.Open(mysql.Open(dbRootURL), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to MySQL server: " + err.Error())
	}

	// Create database if not exists
	sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName)
	if err := dbRoot.Exec(sql).Error; err != nil {
		panic("Failed to create database: " + err.Error())
	}

	// connect to the specific database
	for range 10 {
		DB, err = gorm.Open(mysql.Open(dbURL), &gorm.Config{})
		if err == nil {
			break
		}
		fmt.Println("Waiting for database to be ready...")
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		panic("Failed to connect to database: " + err.Error())
	}

	// migrate models
	if err := DB.AutoMigrate(
		&models.User{},
		&models.Token{},
		&models.Profile{},
		&models.Product{},
		&models.ProductGallery{},
		&models.Category{},
		&models.Address{},
		&models.Province{},
		&models.City{},
		&models.District{},
		&models.Subdistrict{},
		&models.PostalCode{},
		&models.Review{},
		&models.Order{},
		&models.OrderItem{},
		&models.Voucher{},
		&models.UsedVoucher{},
		&models.Notification{},
		&models.NotificationSetting{},
		&models.NotificationType{},
		&models.Payment{},
	); err != nil {
		panic("Migration failed: " + err.Error())
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to get database connection: " + err.Error())
	}
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println("Database connection established successfully.")
}
