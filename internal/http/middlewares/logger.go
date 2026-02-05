package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)


func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		startTime := time.Now()

		
		c.Next()

	
		endTime := time.Now()
		latency := endTime.Sub(startTime)

	
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		
		statusColor := "\033[0m"
		methodColor := "\033[0m" 

		// Status Code Colors
		if statusCode >= 200 && statusCode < 300 {
			statusColor = "\033[32m" 
		} else if statusCode >= 300 && statusCode < 400 {
			statusColor = "\033[33m" 
		} else if statusCode >= 400 && statusCode < 500 {
			statusColor = "\033[34m"
		} else if statusCode >= 500 {
			statusColor = "\033[31m" 
		}

		// Method Colors
		switch reqMethod {
		case "GET":
			methodColor = "\033[34m"
		case "POST":
			methodColor = "\033[32m" 
		case "PUT":
			methodColor = "\033[33m" 
		case "DELETE":
			methodColor = "\033[31m" 
		}

		resetColor := "\033[0m"

		fmt.Printf("[ERP-LOG] %v | %s %3d %s | %13v | %15s | %s %-7s %s | %#v\n",
			endTime.Format("2006/01/02 - 15:04:05"),
			statusColor, statusCode, resetColor,
			latency,
			clientIP,
			methodColor, reqMethod, resetColor,
			reqUri,
		)
	}
}