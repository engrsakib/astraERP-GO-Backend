package routes

import (
	"github.com/engrsakib/erp-system/internal/http/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func RegisterUserRoutes(rg *gin.RouterGroup, db *gorm.DB) {
	
	userHandler := handlers.NewUserHandler(db)

	
	users := rg.Group("/users")
	{
		users.POST("", userHandler.CreateUser)
		users.GET("", userHandler.GetUsers)
		users.GET("/:id", userHandler.GetUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}
}