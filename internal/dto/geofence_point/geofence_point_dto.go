package geofence_point

import "github.com/engrsakib/erp-system/internal/dto/common"

type GeofencePointRequest struct {
    GeofenceID int64   `json:"geofence_id" binding:"required"`
    Latitude   float64 `json:"latitude" binding:"required"`
    Longitude  float64 `json:"longitude" binding:"required"`
    SeqOrder   int     `json:"seq_order"`
}

type GeofencePointSearchQuery struct {
    common.PaginationQuery
    GeofenceID int64  `form:"geofence_id"`
    Search     string `form:"search"`
}
