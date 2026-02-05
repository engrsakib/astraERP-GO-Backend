package middlewares

import (
	"github.com/gin-gonic/gin"
)


func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() 

	
		if len(c.Errors) > 0 {
			
			err := c.Errors.Last()
			
			
			statusCode := c.Writer.Status()
			if statusCode == 200 {
				statusCode = 500
			}

			c.JSON(statusCode, gin.H{
				"status":  "error",
				"message": err.Error(),
			})
		}
	}
}