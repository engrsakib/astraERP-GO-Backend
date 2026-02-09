package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	dto "github.com/engrsakib/erp-system/internal/dto/auth"
)

// RefreshToken godoc
// @Summary      Get New Access Token
// @Description  Exchange a valid refresh token for a new access token and refresh token
// @Tags         Auth
// @Accept       json
// @Produce      json
// @Param        request body auth.RefreshTokenRequest true "Refresh Token"
// @Success      200  {object}  auth.RefreshTokenResponse
// @Failure      400  {object}  map[string]interface{}  "Invalid Request"
// @Failure      401  {object}  map[string]interface{}  "Unauthorized"
// @Router       /api/v1/auth/refresh-token [post]
func (h *AuthHandler) RefreshToken(c *gin.Context) {
	var req dto.RefreshTokenRequest

	// ১. ইনপুট ভ্যালিডেশন
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// ২. সার্ভিস কল
	// ✅ ফিক্স: h.LoginService ব্যবহার করা হয়েছে কারণ AuthHandler এ একাধিক সার্ভিস থাকে
	resp, err := h.LoginService.RefreshToken(req)
	
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	// ৩. রেসপন্স
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   resp,
	})
}