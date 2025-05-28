package main

import (
	"log"
	"os"
	"server/internal/bootstrap"
	"server/internal/config"
	"server/internal/cron"
	"server/internal/middleware"
	"server/internal/routes"
	"server/internal/seeders"
	"server/internal/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.LoadEnv()
	config.InitRedis()
	config.InitMailer()
	config.InitDatabase()
	config.InitMidtrans()
	config.InitCloudinary()
	config.InitGoogleOAuthConfig()
	// config.InitRabbitMQ()

	db := config.DB
	// ========== Seeder ==========
	seeders.ResetDatabase(db)

	// middleware config
	r := gin.Default()
	err := r.SetTrustedProxies(config.GetTrustedProxies())
	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}
	r.Use(
		middleware.Logger(),
		middleware.Recovery(),
		middleware.CORS(),
		middleware.RateLimiter(5, 10),
		middleware.LimitFileSize(12<<20),
		middleware.APIKeyGateway([]string{"/api/payments", "/api/payments/notifications", "/api/auth/google", "/api/auth/google/callback"}),
	)

	// ========== layer ==========
	repo := bootstrap.InitRepositories(db)
	s := bootstrap.InitServices(repo)
	h := bootstrap.InitHandlers(s)
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
