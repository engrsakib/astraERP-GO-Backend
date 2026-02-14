package group_handlers

import (
    "github.com/engrsakib/erp-system/internal/dto/group"
    "github.com/engrsakib/erp-system/internal/services/groupe/group_type"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// @Summary Create Group Type
// @Description Create a new group type
// @Tags GroupType
// @Accept json
// @Produce json
// @Param data body group.GroupTypeRequest true "Group Type Data"
// @Success 201 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /group-types [post]
func CreateGroupTypeHandler(service *group_type.CreateGroupTypeService) gin.HandlerFunc {
    return func(c *gin.Context) {
        var req group.GroupTypeRequest

        if err := c.ShouldBindJSON(&req); err != nil {
            utils.SendError(c, 400, "Invalid request payload", err)
            return
        }

        result, err := service.Execute(req)
        if err != nil {
            utils.SendError(c, 400, "Failed to create group type", err)
            return
        }

        utils.SendResponse(c, 201, "Group type created successfully", result, nil)
    }
}
