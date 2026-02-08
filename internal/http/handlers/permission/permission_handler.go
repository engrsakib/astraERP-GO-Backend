package permission

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/engrsakib/erp-system/internal/dto/user/permission"
	"github.com/engrsakib/erp-system/internal/models"
	permissionService "github.com/engrsakib/erp-system/internal/services/permission"
)

type PermissionHandler struct {
	Service *permissionService.PermissionService
}

func NewPermissionHandler(service *permissionService.PermissionService) *PermissionHandler {
	return &PermissionHandler{Service: service}
}

// AssignPermissions handles the request to update user permissions
func (h *PermissionHandler) AssignPermissions(c *gin.Context) {

	
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No claims found"})
		return
	}

	
	userData, ok := claims.(map[string]interface{})
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token claims format"})
		return
	}

	
	userTypeFloat, ok := userData["user_type"].(float64)
	if !ok {
		c.JSON(http.StatusForbidden, gin.H{"error": "User type missing or invalid in token"})
		return
	}
	requesterType := int(userTypeFloat)


	if requesterType != models.UserTypeSuperAdmin { 
		c.JSON(http.StatusForbidden, gin.H{"error": "Access Denied. Only Super Admin can manage permissions."})
		return
	}

	
	var req dto.AssignPermissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON: " + err.Error()})
		return
	}

	
	if err := h.Service.AssignPermissions(req); err != nil {
		
		fmt.Println("❌ SERVICE ERROR:", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": "Failed to assign permission",
			"details": err.Error(), 
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "User permissions updated successfully",
	})
}