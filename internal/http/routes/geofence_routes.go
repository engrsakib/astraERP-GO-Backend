package routes

import (
    "github.com/gin-gonic/gin"
    "gorm.io/gorm"

    "github.com/engrsakib/erp-system/internal/http/middlewares"
    handler "github.com/engrsakib/erp-system/internal/http/handlers/geofence"
    repo "github.com/engrsakib/erp-system/internal/repository/geofence"
    service "github.com/engrsakib/erp-system/internal/services/geofence"
)

func RegisterGeofenceRoutes(rg *gin.RouterGroup, db *gorm.DB) {

    dao := repo.NewGeofenceRepository(db)

    createService := service.NewCreateGeofenceService(dao)
    getAllService := service.NewGetAllGeofenceService(dao)
    getByIDService := service.NewGetGeofenceByIDService(dao)
    updateService := service.NewUpdateGeofenceService(dao)
    deleteService := service.NewDeleteGeofenceService(dao)

    routes := rg.Group("/geofences")
    routes.Use(middlewares.JWTAuth())
    {
        routes.POST("", middlewares.CheckPermission(db, "geofence.create"), handler.CreateGeofenceHandler(createService))
        routes.GET("", middlewares.CheckPermission(db, "geofence.view"), handler.GetAllGeofencesHandler(getAllService))
        routes.GET("/:id", middlewares.CheckPermission(db, "geofence.view"), handler.GetGeofenceByIDHandler(getByIDService))
        routes.PUT("/:id", middlewares.CheckPermission(db, "geofence.edit"), handler.UpdateGeofenceHandler(updateService))
        routes.DELETE("/:id", middlewares.CheckPermission(db, "geofence.delete"), handler.DeleteGeofenceHandler(deleteService))
    }
}
