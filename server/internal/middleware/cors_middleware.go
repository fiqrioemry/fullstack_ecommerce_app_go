package middleware

import (
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	envOrigins := os.Getenv("ALLOWED_ORIGINS")
	log.Printf("🔍 CORS: ALLOWED_ORIGINS = '%s'", envOrigins)

	allowedOrigins := make(map[string]bool)
	for _, origin := range strings.Split(envOrigins, ",") {
		trimmed := strings.TrimSpace(origin)
		allowedOrigins[trimmed] = true
		log.Printf("✅ CORS: Added allowed origin: '%s'", trimmed)
	}

	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")
		method := c.Request.Method

		log.Printf("🌐 CORS: Request - Method: %s, Origin: '%s'", method, origin)

		// Always set CORS headers first
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Updated headers list - more comprehensive
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, X-API-KEY, Origin, Cache-Control, X-Requested-With, Accept-Encoding, User-Agent, Referer")

		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")

		// Check if origin is allowed
		if allowedOrigins[origin] {
			log.Printf("✅ CORS: Origin ALLOWED: '%s'", origin)
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			log.Printf("❌ CORS: Origin NOT ALLOWED: '%s'", origin)
			log.Printf("🔍 CORS: Available origins: %v", getAllowedOriginsList(allowedOrigins))
		}

		// Handle preflight OPTIONS request
		if method == "OPTIONS" {
			log.Printf("✈️  CORS: Handling OPTIONS preflight for origin: '%s'", origin)
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Helper function to get list of allowed origins for logging
func getAllowedOriginsList(allowedOrigins map[string]bool) []string {
	origins := make([]string, 0, len(allowedOrigins))
	for origin := range allowedOrigins {
		origins = append(origins, origin)
	}
	return origins
}
