package routes

import (
	"github.com/engrsakib/erp-system/internal/http/handlers"
	"github.com/engrsakib/erp-system/internal/http/middlewares"
	"github.com/engrsakib/erp-system/internal/repository/login"
	"github.com/engrsakib/erp-system/internal/services/user"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func RegisterUserRoutes(rg *gin.RouterGroup, db *gorm.DB, redisClient *redis.Client) {
    // User CRUD handler
    userHandler := handlers.NewUserHandler(db)

    users := rg.Group("/users")
    {
        users.POST("", userHandler.CreateUser)
        users.GET("", userHandler.GetUsers)
        users.GET("/:id", userHandler.GetUser)
        users.PUT("/:id", userHandler.UpdateUser)
        users.DELETE("/:id", userHandler.DeleteUser)
    }

   
  

    // OTP handler
    otpService := user.NewOTPService(redisClient)
    userService := user.NewUserService(db)
    
    loginRepo := login.NewLoginRepository(db)
    loginService := user.NewLoginService(loginRepo, redisClient)

    authHandler := handlers.NewAuthHandler(otpService , userService, loginService)

    auth := rg.Group("/auth")
    {
        auth.POST("/send-otp", authHandler.SendOTP)
		auth.POST("/verify-otp", authHandler.VerifyOTP)
        auth.POST("/register", authHandler.RegisterUser)
        auth.POST("/login",middlewares.RateLimiter(redisClient), authHandler.Login)
    }
}
