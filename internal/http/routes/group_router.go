package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "github.com/engrsakib/erp-system/internal/http/middlewares"
    groupHandlers "github.com/engrsakib/erp-system/internal/http/handlers/group"
    groupRepo "github.com/engrsakib/erp-system/internal/repository/group"
    groupService "github.com/engrsakib/erp-system/internal/services/groupe/group_root"
)

func RegisterGroupRoutes(rg *gin.RouterGroup, db *gorm.DB) {

    // Repository
    repo := groupRepo.NewGroupRepository(db)

    // Services
    createService := groupService.NewCreateGroupService(repo)
    getAllService := groupService.NewGetAllGroupService(repo)
    getByIDService := groupService.NewGetGroupByIDService(repo)
    updateService := groupService.NewUpdateGroupService(repo)
    deleteService := groupService.NewDeleteGroupService(repo)

    // Handlers
    createHandler := groupHandlers.CreateGroupHandler(createService)
    getAllHandler := groupHandlers.GetAllGroupHandler(getAllService)
    getByIDHandler := groupHandlers.GetGroupByIDHandler(getByIDService)
    updateHandler := groupHandlers.UpdateGroupHandler(updateService)
    deleteHandler := groupHandlers.DeleteGroupHandler(deleteService)

    groupRoutes := rg.Group("/groups")
    groupRoutes.Use(middlewares.JWTAuth())
    {
        // Create
        groupRoutes.POST("", middlewares.CheckPermission(db, "group.create"), createHandler)

        // Read
        groupRoutes.GET("", middlewares.CheckPermission(db, "group.view"), getAllHandler)
        groupRoutes.GET("/:id", middlewares.CheckPermission(db, "group.view"), getByIDHandler)

        // Update
        groupRoutes.PUT("/:id", middlewares.CheckPermission(db, "group.edit"), updateHandler)

        // Delete
        groupRoutes.DELETE("/:id", middlewares.CheckPermission(db, "group.delete"), deleteHandler)
    }
}
