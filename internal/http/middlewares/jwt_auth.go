package middlewares

import (
	"net/http"


	"github.com/gin-gonic/gin"
	"github.com/engrsakib/erp-system/internal/utils"
)


func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Authorization header is required",
			})
			c.Abort()
			return
		}

		
		claims, err := utils.ExtractTokenPayload(authHeader)
		if err != nil {
			
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Unauthorized: " + err.Error(),
			})
			c.Abort()
			return
		}

		
		c.Set("claims", claims)

		
		c.Next()
	}
}