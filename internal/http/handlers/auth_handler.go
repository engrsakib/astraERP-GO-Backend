package handlers

import (
	"net/http"

	"github.com/engrsakib/erp-system/internal/dto/user/login"
	"github.com/engrsakib/erp-system/internal/services/user"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
    OTPService *user.OTPService
	UserService *user.UserService
    LoginService *user.LoginService
}

func NewAuthHandler(otpService *user.OTPService, userService *user.UserService, loginService *user.LoginService) *AuthHandler {
    return &AuthHandler{
        OTPService:  otpService,
        UserService: userService,
        LoginService: loginService,
    }
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

// VerifyOTP godoc
// @Summary      Verify OTP
// @Description  Verify OTP and return temporary JWT token
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        body  body  struct{Mobile string; OTP string}  true  "Mobile and OTP"
// @Success      200   {object} map[string]interface{}
// @Failure      400   {object} map[string]string
// @Router       /auth/verify-otp [post]
func (h *AuthHandler) VerifyOTP(c *gin.Context) {
    var req struct {
        Mobile string `json:"mobile"`
        OTP    string `json:"otp"`
    }

    if err := c.ShouldBindJSON(&req); err != nil || req.Mobile == "" || req.OTP == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "mobile and otp required"})
        return
    }

    token, err := h.OTPService.VerifyOTP(req.Mobile, req.OTP)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "token":   token,
    })
}


// RegisterUser godoc 
// @Summary Register new user 
// @Description Register a new user using temporary JWT token from OTP verification 
// @Tags auth 
// @Accept json 
// @Produce json 
// @Param Authorization header string true "Temporary JWT Token" 
// @Param body body struct{Name string; Email string; Password string; Confirm string} true "User registration data" 
// @Success 200 {object} map[string]interface{} "User registered successfully" 
// @Failure 400 {object} map[string]string "Invalid input or token" 
// @Router /auth/register [post]
func (h *AuthHandler) RegisterUser(c *gin.Context) {
    var req struct {
        Name     string `json:"name"`
        Email    string `json:"email"`
        Password string `json:"password"`
        Confirm  string `json:"confirm"`
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
        return
    }

    authHeader := c.GetHeader("Authorization")

    err := h.UserService.RegisterUser(authHeader, req.Name, req.Email, req.Password, req.Confirm)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"success": true, "message": "User registered"})
}




// Login godoc 
// @Summary Login user 
// @Description Login using mobile & password and receive access + refresh tokens 
// @Tags auth 
// @Accept json 
// @Produce json 
// @Param body body login.UserLoginRequest true "Login credentials" 
// @Success 200 {object} map[string]interface{} "Login successful" 
// @Failure 400 {object} map[string]string "Invalid input" 
// @Failure 401 {object} map[string]string "Unauthorized"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
    var req login.UserLoginRequest

    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    result, err := h.LoginService.Login(req)
    if err != nil {
        c.JSON(401, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, gin.H{
        "success": true,
        "data":    result,
    })
}