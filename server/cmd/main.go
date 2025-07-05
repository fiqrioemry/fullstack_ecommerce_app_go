package main

import (
	"log"
	"net/http"
	"os"
	"server/internal/config"
	"server/internal/cron"
	"server/internal/handlers"
	"server/internal/middleware"
	"server/internal/repositories"
	"server/internal/routes"
	"server/internal/seeders"
	"server/internal/services"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

// ECOMMERCE APP SERVER
// VERSION: 1.2.0
// DEPLOYMENT: docker-compose
// PORT: 5002
// DESCRIPTION: This is a server for an ecommerce system that handles user registration, product management, and payment processing.

func main() {
	utils.LoadEnv()
	config.InitDependencies()

	db := config.DB
	// ========== Seeder ==========
	seeders.ResetDatabase(db)

	// middleware config
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		log.Println("Root endpoint accessed just now", gin.H{
			"ip":        c.ClientIP(),
			"userAgent": c.GetHeader("User-Agent"),
		})

		c.JSON(http.StatusOK, gin.H{
			"status":    "success",
			"message":   "Welcome to ecommerce API",
			"version":   "1.5.0",
			"ip":        c.ClientIP(),
			"userAgent": c.GetHeader("User-Agent"),
		})
	})

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":    "healthy",
			"message":   "Server is running smoothly",
			"timestamp": utils.NowISO(),
			"uptime":    utils.GetUptime(),
		})
	})

	r.Use(
		middleware.Logger(),
		middleware.Recovery(),
		middleware.CORS(),
		middleware.RateLimiter(5, 10),
		middleware.LimitFileSize(12<<20),
		middleware.APIKeyGateway([]string{"/api/payments", "/api/payments/notifications", "/api/auth/google", "/api/auth/google/callback"}),
	)

	// ========== initialisasi layer ============
	repo := repositories.InitRepositories(db)
	s := services.InitServices(repo)
	h := handlers.InitHandlers(s)

	// ========== Cron Job ==========
	cronManager := cron.NewCronManager(s.PaymentService, s.NotificationService)
	cronManager.RegisterJobs()
	cronManager.Start()

	// ========== Route Binding ==========
	routes.AdminRoutes(r, h.AdminHandler)
	routes.AuthRoutes(r, h.AuthHandler)
	routes.ProfileRoutes(r, h.ProfileHandler)
	routes.BannerRoutes(r, h.BannerHandler)
	routes.CartRoutes(r, h.CartHandler)
	routes.PaymentRoutes(r, h.PaymentHandler)
	routes.ReviewRoutes(r, h.ReviewHandler)
	routes.OrderRoutes(r, h.OrderHandler)
	routes.AddressRoutes(r, h.AddressHandler)
	routes.VoucherRoutes(r, h.VoucherHandler)
	routes.ProductRoutes(r, h.ProductHandler)
	routes.CategoryRoutes(r, h.CategoryHandler)
	routes.LocationRoutes(r, h.LocationHandler)
	routes.NotificationRoutes(r, h.NotificationHandler)

	// ========== Start Server ==========
	port := os.Getenv("PORT")
	log.Println("server running on port:", port)
	log.Fatal(r.Run(":" + port))
}
