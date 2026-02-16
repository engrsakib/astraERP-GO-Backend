package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	_ "github.com/engrsakib/erp-system/internal/docs"
	"github.com/engrsakib/erp-system/internal/http/handlers"
	"github.com/engrsakib/erp-system/internal/http/middlewares"
)

func NewRouter(db *gorm.DB, rdb *redis.Client) *gin.Engine {
	
	r := gin.New()

	r.Use(middlewares.CorsMiddleware())
	r.Use(middlewares.Logger())             
	r.Use(middlewares.GlobalPanicRecovery()) 


	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to ERP System API",
			"status":  "active",
			"docs":    "http://localhost:8080/swagger/index.html",
		})
	})

	
	r.GET("/health", handlers.HealthCheck)


	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	
	v1 := r.Group("/api/v1")
	{
	
		RegisterUserRoutes(v1, db, rdb)
		RegisterPermissionRoutes(v1, db)
		RegisterFaqRoutes(v1, db)
		RegisterGroupTypeRoutes(v1, db)
		RegisterGroupRoutes(v1, db)
		RegisterMemberRoutes(v1, db)
		RegisterGeofenceRoutes(v1, db)
	}

	return r
}