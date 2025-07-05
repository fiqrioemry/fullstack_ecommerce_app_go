package config

func InitDependencies() {
	// Initialize Redis
	InitRedis()

	// Initialize Mailer
	InitMailer()

	// Initialize Database
	InitDatabase()

	// Initialize Cloudinary
	InitCloudinary()

	// Initialize Midtrans
	InitMidtrans()

	// Initialize Google OAuth Config
	InitGoogleOAuthConfig()

}
