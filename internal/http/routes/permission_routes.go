package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	
	permissionHandler "github.com/engrsakib/erp-system/internal/http/handlers/permission"
	"github.com/engrsakib/erp-system/internal/http/middlewares"
	permissionRepo "github.com/engrsakib/erp-system/internal/repository/permission"
	permissionService "github.com/engrsakib/erp-system/internal/services/permission"
)

// RegisterPermissionRoutes: এখানে সব ইনিশিলাইজেশন হবে
func RegisterPermissionRoutes(rg *gin.RouterGroup, db *gorm.DB) {

	// ১. ডিপেন্ডেন্সি তৈরি করা (Repo -> Service -> Handler)
	repo := permissionRepo.NewPermissionRepository(db)
	service := permissionService.NewPermissionService(repo)
	handler := permissionHandler.NewPermissionHandler(service)

	// ২. রাউট গ্রুপ তৈরি করা
	// ধরি, এই রাউটগুলো অ্যাক্সেস করতে হলে লগইন থাকা লাগবে
	permissionGroup := rg.Group("/admin")
	permissionGroup.Use(middlewares.JWTAuth()) // এখানে মিডলওয়্যার বসালাম
	{
		// POST /api/v1/admin/assign-permission
		permissionGroup.POST("/assign-permission", handler.AssignPermissions)
	}
}