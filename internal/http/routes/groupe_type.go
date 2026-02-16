package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "github.com/engrsakib/erp-system/internal/http/middlewares"
    groupHandlers "github.com/engrsakib/erp-system/internal/http/handlers/group_handlers"
    groupRepo "github.com/engrsakib/erp-system/internal/repository/group"
    groupService "github.com/engrsakib/erp-system/internal/services/groupe/group_type"
)

func RegisterGroupTypeRoutes(rg *gin.RouterGroup, db *gorm.DB) {

    // Repository
    repo := groupRepo.NewGroupTypeRepository(db)

    // Services
    createService := groupService.NewCreateGroupTypeService(repo)
    getAllService := groupService.NewGetAllGroupTypeService(repo)
    getByIDService := groupService.NewGetGroupTypeByIDService(repo)
    updateService := groupService.NewUpdateGroupTypeService(repo)
    deleteService := groupService.NewDeleteGroupTypeService(repo)

    // Handlers
    createHandler := groupHandlers.CreateGroupTypeHandler(createService)
    getAllHandler := groupHandlers.GetAllGroupTypeHandler(getAllService)
    getByIDHandler := groupHandlers.GetGroupTypeByIDHandler(getByIDService)
    updateHandler := groupHandlers.UpdateGroupTypeHandler(updateService)
    deleteHandler := groupHandlers.DeleteGroupTypeHandler(deleteService)

    groupTypeGroup := rg.Group("/group-types")
    groupTypeGroup.Use(middlewares.JWTAuth())
    {
        // Create
        groupTypeGroup.POST("", middlewares.CheckPermission(db, "group_type.create"), createHandler)

        // Read
        groupTypeGroup.GET("", middlewares.CheckPermission(db, "group_type.view"), getAllHandler)
        groupTypeGroup.GET("/:id", middlewares.CheckPermission(db, "group_type.view"), getByIDHandler)

        // Update
        groupTypeGroup.PUT("/:id", middlewares.CheckPermission(db, "group_type.edit"), updateHandler)

        // Delete
        groupTypeGroup.DELETE("/:id", middlewares.CheckPermission(db, "group_type.delete"), deleteHandler)
    }
}
