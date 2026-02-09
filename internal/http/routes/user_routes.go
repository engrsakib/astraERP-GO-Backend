package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	
	userHandlers "github.com/engrsakib/erp-system/internal/http/handlers/user"
	"github.com/engrsakib/erp-system/internal/http/middlewares"
	"github.com/engrsakib/erp-system/internal/repository/login"
	userRepo "github.com/engrsakib/erp-system/internal/repository/user" 
	userService "github.com/engrsakib/erp-system/internal/services/user" 
)

func RegisterUserRoutes(rg *gin.RouterGroup, db *gorm.DB, redisClient *redis.Client) {

	
	uRepo := userRepo.NewUserRepository(db)
	lRepo := login.NewLoginRepository(db)

	uService := userService.NewUserService(db, uRepo) 
	lService := userService.NewLoginService(lRepo, redisClient)
	otpService := userService.NewOTPService(redisClient)

	
	userHandler := userHandlers.NewUserHandler(db, uService)

	authHandler := userHandlers.NewAuthHandler(otpService, uService, lService)


	users := rg.Group("/users")
	
	
	users.Use(middlewares.JWTAuth()) 
	{
		// ➤ Create User (Permission: user.create)
		users.POST("", middlewares.CheckPermission(db, "user.create"), userHandler.CreateUser)

		// ➤ Get All Users (Permission: user.view)
		users.GET("/", 
			middlewares.CheckPermission(db, "user.view"), 
			userHandler.GetUsers,
		)

		// ➤ Get Single User (Permission: user.view)
		users.GET("/:id", middlewares.CheckPermission(db, "user.view"), userHandler.GetUser)

		// ➤ Update User (Permission: user.update)
		users.PUT("/:id", middlewares.CheckPermission(db, "user.update"), userHandler.UpdateUser)

		// ➤ Delete User (Permission: user.delete)
		users.DELETE("/:id", middlewares.CheckPermission(db, "user.delete"), userHandler.DeleteUser)
	}

	// ==========================================
	// ৩. Authentication Routes
	// ==========================================
	auth := rg.Group("/auth")
	{
		auth.POST("/send-otp", authHandler.SendOTP)
		auth.POST("/verify-otp", authHandler.VerifyOTP)
		auth.POST("/register", authHandler.RegisterUser)
		
		
		auth.POST("/login", middlewares.RateLimiter(redisClient), authHandler.Login)
		auth.POST("/refresh-token", authHandler.RefreshToken)
	}
}