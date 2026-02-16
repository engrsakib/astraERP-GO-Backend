package group_member

import (
    groupMemberService "github.com/engrsakib/erp-system/internal/services/groupe/group_member"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// GetMemberByIDHandler godoc
// @Summary Get member by ID
// @Description Get a single group member by ID
// @Tags Group Members
// @Param id path int true "Member ID"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /group-members/{id} [get]
func GetMemberByIDHandler(service *groupMemberService.GetMemberByIDService) gin.HandlerFunc {
    return func(c *gin.Context) {

        id, err := utils.ParamID(c)
        if err != nil {
            utils.SendError(c, 400, "Invalid ID", err)
            return
        }

        data, err := service.Execute(int64(id))
        if err != nil {
            utils.SendError(c, 500, "Failed to fetch member", err)
            return
        }

        utils.SendResponse(c, 200, "Member retrieved successfully", data, nil)
    }
}
