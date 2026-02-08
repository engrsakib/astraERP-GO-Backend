package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	userHandlers "github.com/engrsakib/erp-system/internal/http/handlers/user"
	"github.com/engrsakib/erp-system/internal/http/middlewares"
	"github.com/engrsakib/erp-system/internal/repository/login"
	"github.com/engrsakib/erp-system/internal/services/user"
)

func RegisterUserRoutes(rg *gin.RouterGroup, db *gorm.DB, redisClient *redis.Client) {

	userHandler := userHandlers.NewUserHandler(db)

	users := rg.Group("/users")
	
	// users.Use(middlewares.JWTAuth()) 
	{
		users.POST("", userHandler.CreateUser)
		users.GET("", userHandler.GetUsers)
		users.GET("/:id", userHandler.GetUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}

	// ===========================
	// 2. Authentication Routes
	// ===========================

	// Services Initialization
	otpService := user.NewOTPService(redisClient)
	userService := user.NewUserService(db)
	
	loginRepo := login.NewLoginRepository(db)
	loginService := user.NewLoginService(loginRepo, redisClient)

	// 👇 Auth Handler Initialization (userHandlers থেকে কল করা হচ্ছে)
	authHandler := userHandlers.NewAuthHandler(otpService, userService, loginService)

	auth := rg.Group("/auth")
	{
		// OTP Related
		auth.POST("/send-otp", authHandler.SendOTP)
		auth.POST("/verify-otp", authHandler.VerifyOTP)
		
		// Registration & Login
		auth.POST("/register", authHandler.RegisterUser)
		auth.POST("/login", middlewares.RateLimiter(redisClient), authHandler.Login)
		
		auth.POST("/refresh-token", authHandler.RefreshToken)
	}
}