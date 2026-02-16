package geofence

import (
    "github.com/engrsakib/erp-system/internal/dto/geofence"
    service "github.com/engrsakib/erp-system/internal/services/geofence"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// CreateGeofenceHandler godoc
// @Summary Create a new geofence
// @Description Create a new geofence
// @Tags Geofence
// @Accept json
// @Produce json
// @Param geofence body geofence.GeofenceRequest true "Geofence Data"
// @Success 201 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /geofences [post]
func CreateGeofenceHandler(s *service.CreateGeofenceService) gin.HandlerFunc {
    return func(c *gin.Context) {

        var req geofence.GeofenceRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            utils.SendError(c, 400, "Invalid request body", err)
            return
        }

        data, err := s.Execute(req)
        if err != nil {
            utils.SendError(c, 500, "Failed to create geofence", err)
            return
        }

        utils.SendResponse(c, 201, "Geofence created successfully", data, nil)
    }
}
