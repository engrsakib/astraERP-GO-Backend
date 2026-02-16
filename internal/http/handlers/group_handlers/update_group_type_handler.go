package group_handlers

import (
    "strconv"

    "github.com/engrsakib/erp-system/internal/dto/group"
    "github.com/engrsakib/erp-system/internal/services/groupe/group_type"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// @Summary Update Group Type
// @Description Update an existing group type
// @Tags GroupType
// @Accept json
// @Produce json
// @Param id path int true "Group Type ID"
// @Param data body group.GroupTypeRequest true "Group Type Data"
// @Success 200 {object} utils.APIResponse
// @Failure 400 {object} utils.APIResponse
// @Router /group-types/{id} [put]
func UpdateGroupTypeHandler(service *group_type.UpdateGroupTypeService) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

        var req group.GroupTypeRequest
        if err := c.ShouldBindJSON(&req); err != nil {
            utils.SendError(c, 400, "Invalid request payload", err)
            return
        }

        result, err := service.Execute(id, req)
        if err != nil {
            utils.SendError(c, 400, "Failed to update group type", err)
            return
        }

        utils.SendResponse(c, 200, "Group type updated successfully", result, nil)
    }
}
