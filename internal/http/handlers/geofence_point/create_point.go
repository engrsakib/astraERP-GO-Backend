package geofence_point

import (
    "github.com/engrsakib/erp-system/internal/dto/geofence_point"
    service "github.com/engrsakib/erp-system/internal/services/geofence_point"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// CreateGeofencePointHandler godoc
// @Summary Create geofence point
// @Tags Geofence Point
// @Accept json
// @Produce json
// @Param point body geofence_point.GeofencePointRequest true "Point Data"
// @Success 201 {object} utils.APIResponse
// @Router /geofence-points [post]
func CreateGeofencePointHandler(s *service.CreateGeofencePointService) gin.HandlerFunc {
    return func(c *gin.Context) {

        var req geofence_point.GeofencePointRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            utils.SendError(c, 400, "Invalid request body", err)
            return
        }

        data, err := s.Execute(req)
        if err != nil {
            utils.SendError(c, 500, "Failed to create point", err)
            return
        }

        utils.SendResponse(c, 201, "Point created successfully", data, nil)
    }
}
