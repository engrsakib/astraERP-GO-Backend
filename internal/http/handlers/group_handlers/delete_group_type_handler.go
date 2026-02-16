package group_handlers

import (
    "strconv"

    "github.com/engrsakib/erp-system/internal/services/groupe/group_type"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// @Summary Delete Group Type
// @Description Delete a group type by ID
// @Tags GroupType
// @Produce json
// @Param id path int true "Group Type ID"
// @Success 200 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /group-types/{id} [delete]
func DeleteGroupTypeHandler(service *group_type.DeleteGroupTypeService) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

        if err := service.Execute(id); err != nil {
            utils.SendError(c, 404, "Failed to delete group type", err)
            return
        }

        utils.SendResponse(c, 200, "Group type deleted successfully", nil, nil)
    }
}
