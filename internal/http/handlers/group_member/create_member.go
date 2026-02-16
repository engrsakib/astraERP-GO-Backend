package group_member

import (
    "github.com/engrsakib/erp-system/internal/dto/group"
    groupMemberService "github.com/engrsakib/erp-system/internal/services/groupe/group_member"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// CreateMemberHandler godoc
// @Summary Add a member to a group
// @Description Create a new group member
// @Tags Group Members
// @Accept json
// @Produce json
// @Param member body group.MemberRequest true "Member Data"
// @Success 201 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /group-members [post]
func CreateMemberHandler(service *groupMemberService.CreateMemberService) gin.HandlerFunc {
    return func(c *gin.Context) {

        var req group.MemberRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            utils.SendError(c, 400, "Invalid request body", err)
            return
        }

        // Extract user ID from JWT claims
        val, exists := c.Get("claims")
        if !exists {
            utils.SendError(c, 401, "Unauthorized access: user identity not found", nil)
            return
        }

        claims, ok := val.(map[string]interface{})
        if !ok {
            utils.SendError(c, 500, "Invalid claims format", nil)
            return
        }

        var userID int64
        idVal, ok := claims["id"]
        if !ok {
            utils.SendError(c, 401, "Unauthorized: ID not found in claims", nil)
            return
        }

        switch v := idVal.(type) {
        case float64:
            userID = int64(v)
        case int:
            userID = int64(v)
        case int64:
            userID = v
        case uint:
            userID = int64(v)
        default:
            utils.SendError(c, 500, "Unsupported user ID type", nil)
            return
        }

        data, err := service.Execute(req, userID)
        if err != nil {
            utils.SendError(c, 500, "Failed to add member", err)
            return
        }

        utils.SendResponse(c, 201, "Member added successfully", data, nil)
    }
}
