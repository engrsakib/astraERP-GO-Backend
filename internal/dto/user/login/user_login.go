package login

type UserLoginRequest struct {
    Mobile   string `json:"mobile" binding:"required"`
    Password string `json:"password" binding:"required"`
}
