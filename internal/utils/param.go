package utils

import (
	"strconv"
	"github.com/gin-gonic/gin"
)


func ParamID(c *gin.Context) (uint, error) {
	idStr := c.Param("id") 
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}