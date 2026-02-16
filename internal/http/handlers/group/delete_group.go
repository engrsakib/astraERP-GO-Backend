package group

import (
    group_root "github.com/engrsakib/erp-system/internal/services/groupe/group_root"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// DeleteGroupHandler godoc
// @Summary Delete a group
// @Description Delete a group by ID
// @Tags Group
// @Param id path int true "Group ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /groups/{id} [delete]
func DeleteGroupHandler(service *group_root.DeleteGroupService) gin.HandlerFunc {
    return func(c *gin.Context) {

        id, err := utils.ParamID(c)
        if err != nil {
            utils.SendError(c, 400, "Invalid ID", err)
            return
        }

        err = service.Execute(int64(id))
        if err != nil {
            utils.SendError(c, 500, "Failed to delete group", err)
            return
        }

        utils.SendResponse(c, 200, "Group deleted successfully", nil, nil)
    }
}
