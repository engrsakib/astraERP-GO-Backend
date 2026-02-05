package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)


func GlobalPanicRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
			
				log.Printf("🔥 PANIC RECOVERED: %v\n%s", err, debug.Stack())

			
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"status":  "error",
					"message": "Internal Server Error. Something went wrong!",
					"error":   err, 
				})
			}
		}()
		c.Next()
	}
}