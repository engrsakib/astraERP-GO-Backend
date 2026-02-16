package geofence

import (
    "github.com/engrsakib/erp-system/internal/dto/geofence"
    service "github.com/engrsakib/erp-system/internal/services/geofence"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// GetAllGeofencesHandler godoc
// @Summary Get all geofences
// @Description Get all geofences with pagination & search
// @Tags Geofence
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Param search query string false "Search by name"
// @Success 200 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /geofences [get]
func GetAllGeofencesHandler(s *service.GetAllGeofenceService) gin.HandlerFunc {
    return func(c *gin.Context) {

        var q geofence.GeofenceSearchQuery
        _ = c.ShouldBindQuery(&q)

        data, meta, err := s.Execute(q)
        if err != nil {
            utils.SendError(c, 500, "Failed to fetch geofences", err)
            return
        }

        utils.SendResponse(c, 200, "Geofences retrieved successfully", data, meta)
    }
}
