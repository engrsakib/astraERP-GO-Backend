package geofence_point

import (
    service "github.com/engrsakib/erp-system/internal/services/geofence_point"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// DeleteGeofencePointHandler godoc
// @Summary Delete geofence point
// @Tags Geofence Point
// @Param id path int true "Point ID"
// @Success 200 {object} utils.APIResponse
// @Router /geofence-points/{id} [delete]
func DeleteGeofencePointHandler(s *service.DeleteGeofencePointService) gin.HandlerFunc {
    return func(c *gin.Context) {

        id, err := utils.ParamID(c)
        if err != nil {
            utils.SendError(c, 400, "Invalid ID", err)
            return
        }

        err = s.Execute(int64(id))
        if err != nil {
            utils.SendError(c, 500, "Failed to delete point", err)
            return
        }

        utils.SendResponse(c, 200, "Point deleted successfully", nil, nil)
    }
}
