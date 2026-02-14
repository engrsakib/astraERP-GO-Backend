package utils

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

type APIResponse struct {
	Success bool        `json:"success"`
	// Code    int         `json:"code,omitempty"`  
	Message string      `json:"message"`
	Meta    interface{} `json:"meta,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"` 
}


func SendResponse(c *gin.Context, statusCode int, message string, data interface{}, meta interface{}) {
	response := APIResponse{
		Success: statusCode >= 200 && statusCode < 300, 
		// Code:    statusCode,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
	c.JSON(statusCode, response)
}


func SendError(c *gin.Context, status int, message string, err error) {
	
	errorText := ""
	if err != nil {
		errorText = err.Error()
		
		fmt.Printf("❌ [API ERROR] Path: %s | Error: %v\n", c.Request.URL.Path, errorText)
	}

	c.AbortWithStatusJSON(status, APIResponse{
		Success: false,
		// Code:    status,   
		Message: message,
		Error:   errorText, 
	})
}