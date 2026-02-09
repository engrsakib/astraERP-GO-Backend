package user

import (
    userService "github.com/engrsakib/erp-system/internal/services/user"
)


type LoginHandler struct {
    Service *userService.LoginService
}


func NewLoginHandler(service *userService.LoginService) *LoginHandler {
    return &LoginHandler{
        Service: service,
    }
}