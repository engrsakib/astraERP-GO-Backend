package group_member

import (
    groupMemberService "github.com/engrsakib/erp-system/internal/services/groupe/group_member"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// DeleteMemberHandler godoc
// @Summary Delete a group member
// @Description Soft delete a group member by ID
// @Tags Group Members
// @Param id path int true "Member ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /group-members/{id} [delete]
func DeleteMemberHandler(service *groupMemberService.DeleteMemberService) gin.HandlerFunc {
    return func(c *gin.Context) {

        id, err := utils.ParamID(c)
        if err != nil {
            utils.SendError(c, 400, "Invalid ID", err)
            return
        }

        err = service.Execute(int64(id))
        if err != nil {
            utils.SendError(c, 500, "Failed to delete member", err)
            return
        }

        utils.SendResponse(c, 200, "Member deleted successfully", nil, nil)
    }
}
