package group_member

import (
    "github.com/engrsakib/erp-system/internal/dto/group"
    groupMemberService "github.com/engrsakib/erp-system/internal/services/groupe/group_member"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// UpdateMemberHandler godoc
// @Summary Update a group member
// @Description Update member details
// @Tags Group Members
// @Accept json
// @Produce json
// @Param id path int true "Member ID"
// @Param member body group.MemberRequest true "Member Data"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /group-members/{id} [put]
func UpdateMemberHandler(service *groupMemberService.UpdateMemberService) gin.HandlerFunc {
    return func(c *gin.Context) {

        id, err := utils.ParamID(c)
        if err != nil {
            utils.SendError(c, 400, "Invalid ID", err)
            return
        }

        var req group.MemberRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            utils.SendError(c, 400, "Invalid request body", err)
            return
        }

        err = service.Execute(int64(id), req)
        if err != nil {
            utils.SendError(c, 500, "Failed to update member", err)
            return
        }

        utils.SendResponse(c, 200, "Member updated successfully", nil, nil)
    }
}
