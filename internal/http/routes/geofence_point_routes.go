package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "github.com/engrsakib/erp-system/internal/http/middlewares"
    handler "github.com/engrsakib/erp-system/internal/http/handlers/geofence_point"
    repo "github.com/engrsakib/erp-system/internal/repository/geofence_point"
    service "github.com/engrsakib/erp-system/internal/services/geofence_point"
)

func RegisterGeofencePointRoutes(rg *gin.RouterGroup, db *gorm.DB) {

    dao := repo.NewGeofencePointRepository(db)

    createService := service.NewCreateGeofencePointService(dao)
    getAllService := service.NewGetAllGeofencePointService(dao)
    getByIDService := service.NewGetGeofencePointByIDService(dao)
    updateService := service.NewUpdateGeofencePointService(dao)
    deleteService := service.NewDeleteGeofencePointService(dao)

    routes := rg.Group("/geofence-points")
    routes.Use(middlewares.JWTAuth())
    {
        routes.POST("", middlewares.CheckPermission(db, "geofence_point.create"), handler.CreateGeofencePointHandler(createService))
        routes.GET("", middlewares.CheckPermission(db, "geofence_point.view"), handler.GetAllGeofencePointsHandler(getAllService))
        routes.GET("/:id", middlewares.CheckPermission(db, "geofence_point.view"), handler.GetGeofencePointByIDHandler(getByIDService))
        routes.PUT("/:id", middlewares.CheckPermission(db, "geofence_point.edit"), handler.UpdateGeofencePointHandler(updateService))
        routes.DELETE("/:id", middlewares.CheckPermission(db, "geofence_point.delete"), handler.DeleteGeofencePointHandler(deleteService))
    }
}
