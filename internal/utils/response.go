package utils

import (
	"github.com/gin-gonic/gin"
)

// APIResponse: সব রেসপন্সের সাধারণ স্ট্রাকচার
type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"` 
}


func SendResponse(c *gin.Context, statusCode int, message string, data interface{}, meta interface{}) {
	response := APIResponse{
		Success: statusCode >= 200 && statusCode < 300, 
		Message: message,
		Data:    data,
		Meta:    meta,
	}
	c.JSON(statusCode, response)
}


func SendError(c *gin.Context, statusCode int, message string, err error) {
	errMsg := message
	if err != nil {
		
		errMsg = message + ": " + err.Error() 
	}
	
	SendResponse(c, statusCode, errMsg, nil, nil)
}