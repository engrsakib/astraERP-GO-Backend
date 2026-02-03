package routes

import (
    "github.com/engrsakib/erp-system/internal/http/handlers"
    "github.com/engrsakib/erp-system/internal/services/user"
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
    "github.com/redis/go-redis/v9"
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
    authHandler := handlers.NewAuthHandler(otpService)

    auth := rg.Group("/auth")
    {
        auth.POST("/send-otp", authHandler.SendOTP)
		auth.POST("/verify-otp", authHandler.VerifyOTP)
    }
}
