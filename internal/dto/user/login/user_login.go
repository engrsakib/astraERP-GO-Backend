package login

// SendOTPRequest - OTP পাঠানোর জন্য
type SendOTPRequest struct {
    Mobile string `json:"mobile" binding:"required"`
}



// UserLoginRequest - লগইনের জন্য
type UserLoginRequest struct {
    Mobile   string `json:"mobile" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type VerifyOTPRequest struct {
    Mobile string `json:"mobile" binding:"required"`
    OTP    string `json:"otp" binding:"required"`
}

type RegisterRequest struct {
    Name     string `json:"name" binding:"required"`
    Mobile   string `json:"mobile" binding:"required"`
    Password string `json:"password" binding:"required"`
}