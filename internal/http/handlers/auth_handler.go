package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/engrsakib/erp-system/internal/services/user"
)

type AuthHandler struct {
    OTPService *user.OTPService
}

func NewAuthHandler(otpService *user.OTPService) *AuthHandler {
    return &AuthHandler{OTPService: otpService}
}

// SendOTP godoc
// @Summary      Send OTP
// @Description  Send 6-digit OTP to mobile number via SMS
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  struct{Mobile string}  true  "Mobile number"
// @Success      200   {object} map[string]string
// @Failure      400   {object} map[string]string
// @Failure      500   {object} map[string]string
// @Router       /auth/send-otp [post]
func (h *AuthHandler) SendOTP(c *gin.Context) {
    var req struct {
        Mobile string `json:"mobile"`
    }

    if err := c.ShouldBindJSON(&req); err != nil || req.Mobile == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "mobile required"})
        return
    }

    err := h.OTPService.SendOTP(req.Mobile)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "OTP sent"})
}
