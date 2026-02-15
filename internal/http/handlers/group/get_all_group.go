package group

import (
    "github.com/engrsakib/erp-system/internal/dto/common"
    group_root "github.com/engrsakib/erp-system/internal/services/groupe/group_root"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// GetAllGroupHandler godoc
// @Summary Get all groups
// @Description Get all groups with pagination & search
// @Tags Group
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Param search query string false "Search by name"
// @Success 200 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /groups [get]
func GetAllGroupHandler(service *group_root.GetAllGroupService) gin.HandlerFunc {
    return func(c *gin.Context) {

        var q common.PaginationQuery
        _ = c.ShouldBindQuery(&q)

        data, meta, err := service.Execute(q)
        if err != nil {
            utils.SendError(c, 500, "Failed to fetch groups", err)
            return
        }

        utils.SendResponse(c, 200, "Groups retrieved successfully", data, meta)
    }
}
