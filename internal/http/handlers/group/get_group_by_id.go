package group

import (
    group_root "github.com/engrsakib/erp-system/internal/services/groupe/group_root"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// GetGroupByIDHandler godoc
// @Summary Get group by ID
// @Description Get a single group by ID
// @Tags Group
// @Param id path int true "Group ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /groups/{id} [get]
func GetGroupByIDHandler(service *group_root.GetGroupByIDService) gin.HandlerFunc {
    return func(c *gin.Context) {

        id, err := utils.ParamID(c)
        if err != nil {
            utils.SendError(c, 400, "Invalid ID", err)
            return
        }

        data, err := service.Execute(int64(id))
        if err != nil {
            utils.SendError(c, 500, "Failed to fetch group", err)
            return
        }

        utils.SendResponse(c, 200, "Group retrieved successfully", data, nil)
    }
}
