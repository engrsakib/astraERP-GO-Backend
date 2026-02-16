package auth

type LoginRequest struct {
    Mobile string `json:"mobile" binding:"required"`
}