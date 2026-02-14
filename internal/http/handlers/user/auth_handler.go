package user

import (
	"net/http"
	"os"
	"strconv"

	dto "github.com/engrsakib/erp-system/internal/dto/user/login"
	"github.com/engrsakib/erp-system/internal/services/user"
	"github.com/gin-gonic/gin"
    "github.com/engrsakib/erp-system/internal/utils"
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


func (h *AuthHandler) setAuthCookies(c *gin.Context, accessToken, refreshToken string) {
	
	cookieDomain := os.Getenv("COOKIE_DOMAIN")
	if cookieDomain == "" {
		cookieDomain = "localhost"
	}

	cookieSecure := os.Getenv("COOKIE_SECURE") == "true"

	
	accExp, _ := strconv.Atoi(os.Getenv("ACCESS_TOKEN_EXP"))
	refExp, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXP"))

	
	if accExp == 0 { accExp = 60 }        
	if refExp == 0 { refExp = 30240 }    

	// Access Token Cookie
	c.SetCookie("access_token", accessToken, accExp*60, "/", cookieDomain, cookieSecure, true)

	// Refresh Token Cookie
	c.SetCookie("refresh_token", refreshToken, refExp*60, "/", cookieDomain, cookieSecure, true)
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
// @Summary Register a new user
// @Description Register a new user using the temporary JWT token received after OTP verification. Returns Access & Refresh tokens.
// @Tags Auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <Temporary_Token>"
// @Param request body struct{Name string "User Name"; Email string "User Email"; Password string "Password"; Confirm string "Confirm Password"} true "User Registration Data"
// @Success 201 {object} utils.APIResponse{data=map[string]string} "Registration Successful"
// @Failure 400 {object} utils.APIResponse "Invalid Input or Token"
// @Failure 500 {object} utils.APIResponse "Internal Server Error"
// @Router /auth/register [post]
func (h *AuthHandler) RegisterUser(c *gin.Context) {
	var req struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
		Confirm  string `json:"confirm" binding:"required,eqfield=Password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid input data", err)
		return
	}

	authHeader := c.GetHeader("Authorization")
	accessToken, refreshToken, err := h.UserService.RegisterUser(authHeader, req.Name, req.Email, req.Password, req.Confirm)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "Registration failed", err)
		return
	}

	h.setAuthCookies(c, accessToken, refreshToken)

	utils.SendResponse(c, http.StatusCreated, "User registered successfully", gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	}, nil)
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
	var req dto.UserLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.LoginService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	// ✅ হেল্পার মেথড ব্যবহার
	h.setAuthCookies(c, result.AccessToken, result.RefreshToken)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    result,
	})
}