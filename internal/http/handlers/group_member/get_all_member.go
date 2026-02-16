package group_member

import (
    "github.com/engrsakib/erp-system/internal/dto/group"
    groupMemberService "github.com/engrsakib/erp-system/internal/services/groupe/group_member"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// GetAllMembersHandler godoc
// @Summary Get all group members
// @Description Get all members with pagination & search
// @Tags Group Members
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Param search query string false "Search by user name"
// @Success 200 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /group-members [get]
func GetAllMembersHandler(service *groupMemberService.GetAllMemberService) gin.HandlerFunc {
    return func(c *gin.Context) {

        var q group.MemberSearchQuery
        _ = c.ShouldBindQuery(&q)

        data, meta, err := service.Execute(q)
        if err != nil {
            utils.SendError(c, 500, "Failed to fetch members", err)
            return
        }

        utils.SendResponse(c, 200, "Members retrieved successfully", data, meta)
    }
}
