package geofence

import (
    service "github.com/engrsakib/erp-system/internal/services/geofence"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// GetGeofenceByIDHandler godoc
// @Summary Get geofence by ID
// @Description Get a single geofence by ID
// @Tags Geofence
// @Param id path int true "Geofence ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /geofences/{id} [get]
func GetGeofenceByIDHandler(s *service.GetGeofenceByIDService) gin.HandlerFunc {
    return func(c *gin.Context) {

        id, err := utils.ParamID(c)
        if err != nil {
            utils.SendError(c, 400, "Invalid ID", err)
            return
        }

        data, err := s.Execute(int64(id))
        if err != nil {
            utils.SendError(c, 404, "Geofence not found", err)
            return
        }

        utils.SendResponse(c, 200, "Geofence retrieved successfully", data, nil)
    }
}
