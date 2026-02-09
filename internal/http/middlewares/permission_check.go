package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/engrsakib/erp-system/internal/models"
	"gorm.io/gorm"
)


func CheckPermission(db *gorm.DB, requiredSlug string) gin.HandlerFunc {
	return func(c *gin.Context) {
		
		
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: No token claims found"})
			c.Abort()
			return
		}

		
		userData, ok := claims.(map[string]interface{})
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid token format"})
			c.Abort()
			return
		}

		userID := int64(userData["id"].(float64))
		
		
		var userType int
		if ut, ok := userData["user_type"].(float64); ok {
			userType = int(ut)
		}

		if userType == 0 { 
			c.Next()
			return
		}

		var count int64
		err := db.Model(&models.UserPermission{}).
			Where("user_id = ? AND permission_slug = ?", userID, requiredSlug).
			Count(&count).Error

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error while checking permissions"})
			c.Abort()
			return
		}

		if count == 0 {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "Permission Denied! You do not have access to: " + requiredSlug,
			})
			c.Abort() 
			return
		}

		c.Next()
	}
}