package geofence

import (
    "github.com/engrsakib/erp-system/internal/dto/geofence"
    service "github.com/engrsakib/erp-system/internal/services/geofence"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// UpdateGeofenceHandler godoc
// @Summary Update geofence
// @Description Update geofence details
// @Tags Geofence
// @Param id path int true "Geofence ID"
// @Param geofence body geofence.GeofenceRequest true "Geofence Data"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /geofences/{id} [put]
func UpdateGeofenceHandler(s *service.UpdateGeofenceService) gin.HandlerFunc {
    return func(c *gin.Context) {

        id, err := utils.ParamID(c)
        if err != nil {
            utils.SendError(c, 400, "Invalid ID", err)
            return
        }

        var req geofence.GeofenceRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            utils.SendError(c, 400, "Invalid request body", err)
            return
        }

        err = s.Execute(int64(id), req)
        if err != nil {
            utils.SendError(c, 500, "Failed to update geofence", err)
            return
        }

        utils.SendResponse(c, 200, "Geofence updated successfully", nil, nil)
    }
}
