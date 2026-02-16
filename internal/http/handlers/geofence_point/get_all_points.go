package geofence_point

import (
    "github.com/engrsakib/erp-system/internal/dto/geofence_point"
    service "github.com/engrsakib/erp-system/internal/services/geofence_point"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// GetAllGeofencePointsHandler godoc
// @Summary Get all geofence points
// @Tags Geofence Point
// @Param geofence_id query int false "Filter by geofence"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} utils.APIResponse
// @Router /geofence-points [get]
func GetAllGeofencePointsHandler(s *service.GetAllGeofencePointService) gin.HandlerFunc {
    return func(c *gin.Context) {

        var q geofence_point.GeofencePointSearchQuery
        _ = c.ShouldBindQuery(&q)

        data, meta, err := s.Execute(q)
        if err != nil {
            utils.SendError(c, 500, "Failed to fetch points", err)
            return
        }

        utils.SendResponse(c, 200, "Points retrieved successfully", data, meta)
    }
}
