package user

import (
	"net/http"

	"github.com/engrsakib/erp-system/internal/dto/user/login"
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
    // ১. ইনপুট বাইন্ডিং (Validation সহ)
    var req struct {
        Name     string `json:"name" binding:"required"`
        Email    string `json:"email" binding:"required,email"`
        Password string `json:"password" binding:"required,min=6"`
        Confirm  string `json:"confirm" binding:"required,eqfield=Password"` // পাসওয়ার্ড ম্যাচ করছে কিনা চেক করবে
    }

    if err := c.ShouldBindJSON(&req); err != nil {
        utils.SendError(c, http.StatusBadRequest, "Invalid input data. Please check required fields.", err)
        return
    }

    // ২. হেডার থেকে টোকেন নেওয়া
    authHeader := c.GetHeader("Authorization")
    if authHeader == "" {
        utils.SendError(c, http.StatusBadRequest, "Authorization header is missing", nil)
        return
    }

    // ৩. সার্ভিস কল করা (যা এখন ৩টি ভ্যালু রিটার্ন করে)
    accessToken, refreshToken, err := h.UserService.RegisterUser(authHeader, req.Name, req.Email, req.Password, req.Confirm)
    if err != nil {
        // সার্ভিসের এরর ক্লায়েন্টকে পাঠানো
        utils.SendError(c, http.StatusBadRequest, "Registration failed", err)
        return
    }

    // ৪. রেসপন্স ডাটা তৈরি
    responseData := map[string]string{
        "access_token":  accessToken,
        "refresh_token": refreshToken,
    }

    // ৫. সফল রেসপন্স পাঠানো (201 Created)
    utils.SendResponse(c, http.StatusCreated, "User registered successfully", responseData, nil)
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