package middlewares

import (
	"os"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CorsMiddleware() gin.HandlerFunc {
	
	originsString := os.Getenv("CORS_ORIGINS")

	
	var allowedOrigins []string
	if originsString != "" {
		for _, origin := range strings.Split(originsString, ",") {
			allowedOrigins = append(allowedOrigins, strings.TrimSpace(origin))
		}
	}

	// ৩. কনফিগারেশন সেট করা
	config := cors.Config{
		AllowOrigins:     allowedOrigins, 
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, 
		MaxAge:           12 * time.Hour,
	}

	return cors.New(config)
}