package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	permissionHandler "github.com/engrsakib/erp-system/internal/http/handlers/permission"
	"github.com/engrsakib/erp-system/internal/http/middlewares"
	permissionRepo "github.com/engrsakib/erp-system/internal/repository/permission"
	permissionService "github.com/engrsakib/erp-system/internal/services/permission"
)


func RegisterPermissionRoutes(rg *gin.RouterGroup, db *gorm.DB) {

	repo := permissionRepo.NewPermissionRepository(db)
	service := permissionService.NewPermissionService(repo)
	handler := permissionHandler.NewPermissionHandler(service)

	
	permissionGroup := rg.Group("/admin")
	permissionGroup.Use(middlewares.JWTAuth()) 
	{
		// POST /api/v1/admin/assign-permission
		permissionGroup.POST("/assign-permission",middlewares.CheckPermission(db, "permission.assign"), handler.AssignPermissions)
	}
}