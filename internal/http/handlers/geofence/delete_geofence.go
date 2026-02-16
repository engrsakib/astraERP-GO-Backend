package geofence

import (
    service "github.com/engrsakib/erp-system/internal/services/geofence"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// DeleteGeofenceHandler godoc
// @Summary Delete geofence
// @Description Soft delete geofence by ID
// @Tags Geofence
// @Param id path int true "Geofence ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /geofences/{id} [delete]
func DeleteGeofenceHandler(s *service.DeleteGeofenceService) gin.HandlerFunc {
    return func(c *gin.Context) {

        id, err := utils.ParamID(c)
        if err != nil {
            utils.SendError(c, 400, "Invalid ID", err)
            return
        }

        err = s.Execute(int64(id))
        if err != nil {
            utils.SendError(c, 500, "Failed to delete geofence", err)
            return
        }

        utils.SendResponse(c, 200, "Geofence deleted successfully", nil, nil)
    }
}
