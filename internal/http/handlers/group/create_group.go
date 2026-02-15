package group

import (
    "github.com/engrsakib/erp-system/internal/dto/group"
    group_root "github.com/engrsakib/erp-system/internal/services/groupe/group_root"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// CreateGroupHandler godoc
// @Summary Create a new group
// @Description Create a new group
// @Tags Group
// @Accept json
// @Produce json
// @Param group body group.GroupRequest true "Group Data"
// @Success 201 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Failure 500 {object} utils.APIResponse
// @Router /groups [post]
func CreateGroupHandler(service *group_root.CreateGroupService) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		var req group.GroupRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			utils.SendError(c, 400, "Invalid request body", err)
			return
		}

	
		val, exists := c.Get("claims")
		if !exists {
			utils.SendError(c, 401, "Unauthorized access: user identity not found", nil)
			return
		}

		claims, ok := val.(map[string]interface{})
		if !ok {
			utils.SendError(c, 500, "Internal server error: invalid claims format", nil)
			return
		}

		
		var userID int64
		idVal, ok := claims["id"]
		if !ok {
			utils.SendError(c, 401, "Unauthorized access: ID not found in claims", nil)
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
			utils.SendError(c, 500, "Internal server error: unsupported user ID type", nil)
			return
		}

		data, err := service.Execute(req, userID)
		if err != nil {
			utils.SendError(c, 500, "Failed to create group", err)
			return
		}

		
		utils.SendResponse(c, 201, "Group created successfully", data, nil)
	}
}
