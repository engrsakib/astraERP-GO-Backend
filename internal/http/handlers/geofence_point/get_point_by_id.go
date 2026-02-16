package geofence_point

import (
    service "github.com/engrsakib/erp-system/internal/services/geofence_point"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// GetGeofencePointByIDHandler godoc
// @Summary Get point by ID
// @Tags Geofence Point
// @Param id path int true "Point ID"
// @Success 200 {object} utils.APIResponse
// @Router /geofence-points/{id} [get]
func GetGeofencePointByIDHandler(s *service.GetGeofencePointByIDService) gin.HandlerFunc {
    return func(c *gin.Context) {

        id, err := utils.ParamID(c)
        if err != nil {
            utils.SendError(c, 400, "Invalid ID", err)
            return
        }

        data, err := s.Execute(int64(id))
        if err != nil {
            utils.SendError(c, 404, "Point not found", err)
            return
        }

        utils.SendResponse(c, 200, "Point retrieved successfully", data, nil)
    }
}
