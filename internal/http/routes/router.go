package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9" // Redis ইমপোর্ট করতে হবে
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	_ "github.com/engrsakib/erp-system/internal/docs"
	"github.com/engrsakib/erp-system/internal/http/handlers"
)

// ফাংশন সিগনেচার আপডেট করা হয়েছে: (db *gorm.DB, rdb *redis.Client)
func NewRouter(db *gorm.DB, rdb *redis.Client) *gin.Engine {
	r := gin.Default()

	r.GET("/health", handlers.HealthCheck)

	// বর্তমানে UserHandler শুধু DB নিচ্ছে।
	// ভবিষ্যতে যদি হ্যান্ডলারের ভেতর রেডিস লাগে, তখন handlers.NewUserHandler(db, rdb) করতে হবে।
	userHandler := handlers.NewUserHandler(db)

	users := r.Group("/users")
	{
		users.POST("", userHandler.CreateUser)
		users.GET("", userHandler.GetUsers)
		users.GET("/:id", userHandler.GetUser)
		users.PUT("/:id", userHandler.UpdateUser)
		users.DELETE("/:id", userHandler.DeleteUser)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}