package group_handlers

import (
    "github.com/engrsakib/erp-system/internal/services/groupe/group_type"
    "github.com/engrsakib/erp-system/internal/utils"
    "github.com/gin-gonic/gin"
)

// @Summary Get All Group Types
// @Description Retrieve all group types
// @Tags GroupType
// @Produce json
// @Success 200 {object} utils.APIResponse
// @Router /group-types [get]
func GetAllGroupTypeHandler(service *group_type.GetAllGroupTypeService) gin.HandlerFunc {
    return func(c *gin.Context) {
        result, err := service.Execute()
        if err != nil {
            utils.SendError(c, 500, "Failed to fetch group types", err)
            return
        }

        utils.SendResponse(c, 200, "Group types retrieved successfully", result, nil)
    }
}
