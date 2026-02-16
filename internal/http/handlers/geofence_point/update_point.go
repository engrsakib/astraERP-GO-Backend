package geofence_point

import (
    "github.com/engrsakib/erp-system/internal/dto/geofence_point"
    service "github.com/engrsakib/erp-system/internal/services/geofence_point"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// UpdateGeofencePointHandler godoc
// @Summary Update geofence point
// @Tags Geofence Point
// @Param id path int true "Point ID"
// @Param point body geofence_point.GeofencePointRequest true "Point Data"
// @Success 200 {object} utils.APIResponse
// @Router /geofence-points/{id} [put]
func UpdateGeofencePointHandler(s *service.UpdateGeofencePointService) gin.HandlerFunc {
    return func(c *gin.Context) {

        id, err := utils.ParamID(c)
        if err != nil {
            utils.SendError(c, 400, "Invalid ID", err)
            return
        }

        var req geofence_point.GeofencePointRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            utils.SendError(c, 400, "Invalid request body", err)
            return
        }

        err = s.Execute(int64(id), req)
        if err != nil {
            utils.SendError(c, 500, "Failed to update point", err)
            return
        }

        utils.SendResponse(c, 200, "Point updated successfully", nil, nil)
    }
}
