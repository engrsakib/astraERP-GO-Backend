package group_handlers

import (
    "strconv"

    "github.com/engrsakib/erp-system/internal/services/groupe/group_type"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// @Summary Get Group Type by ID
// @Description Retrieve a single group type by ID
// @Tags GroupType
// @Produce json
// @Param id path int true "Group Type ID"
// @Success 200 {object} utils.APIResponse
// @Failure 404 {object} utils.APIResponse
// @Router /group-types/{id} [get]
func GetGroupTypeByIDHandler(service *group_type.GetGroupTypeByIDService) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

        result, err := service.Execute(id)
        if err != nil {
            utils.SendError(c, 404, "Group type not found", err)
            return
        }

        utils.SendResponse(c, 200, "Group type retrieved successfully", result, nil)
    }
}
