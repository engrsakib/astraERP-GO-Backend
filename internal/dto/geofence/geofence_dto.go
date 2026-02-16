package geofence

import "github.com/engrsakib/erp-system/internal/dto/common"

// Create / Update DTO
type GeofenceRequest struct {
    Name        string `json:"name" binding:"required"`
    Description string `json:"description"`
}

// Search + Pagination DTO
type GeofenceSearchQuery struct {
    common.PaginationQuery
    Search string `form:"search"`
}
