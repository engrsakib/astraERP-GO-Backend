package group

import (
    "github.com/engrsakib/erp-system/internal/dto/group"
    group_root "github.com/engrsakib/erp-system/internal/services/groupe/group_root"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// UpdateGroupHandler godoc
// @Summary Update a group
// @Description Update group by ID
// @Tags Group
// @Accept json
// @Produce json
// @Param id path int true "Group ID"
// @Param group body group.GroupRequest true "Group Data"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /groups/{id} [put]
func UpdateGroupHandler(service *group_root.UpdateGroupService) gin.HandlerFunc {
    return func(c *gin.Context) {

        id, err := utils.ParamID(c)
        if err != nil {
            utils.SendError(c, 400, "Invalid ID", err)
            return
        }

        var req group.GroupRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            utils.SendError(c, 400, "Invalid request body", err)
            return
        }

        err = service.Execute(int64(id), req)
        if err != nil {
            utils.SendError(c, 500, "Failed to update group", err)
            return
        }

        utils.SendResponse(c, 200, "Group updated successfully", nil, nil)
    }
}
