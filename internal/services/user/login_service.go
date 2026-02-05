package user

import (
    "errors"

    dto "github.com/engrsakib/erp-system/internal/dto/user/login"
    loginRepo "github.com/engrsakib/erp-system/internal/repository/login"
    "github.com/engrsakib/erp-system/internal/utils"

    "golang.org/x/crypto/bcrypt"
)

type LoginService struct {
    Repo *loginRepo.LoginRepository
}

func NewLoginService(repo *loginRepo.LoginRepository) *LoginService {
    return &LoginService{Repo: repo}
}

func (s *LoginService) Login(req dto.UserLoginRequest) (map[string]interface{}, error) {
    user, err := s.Repo.FindByMobile(req.Mobile)
    if err != nil {
        return nil, errors.New("invalid mobile or password")
    }

    
    if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
        return nil, errors.New("invalid mobile or password")
    }

    // Token payload
    payload := map[string]interface{}{
        "id":     user.ID,
        "name":   user.Name,
        "mobile": user.Mobile,
        "photo":  user.Photo,
    }

    // Access Token
    accessToken, err := utils.GenerateToken(payload, "ACCESS_TOKEN_EXP")
    if err != nil {
        return nil, err
    }

    // Refresh Token
    refreshToken, err := utils.GenerateToken(payload, "REFRESH_TOKEN_EXP")
    if err != nil {
        return nil, err
    }

    return map[string]interface{}{
        "accessToken":  accessToken,
        "refreshToken": refreshToken,
        "user": map[string]interface{}{
            "id":     user.ID,
            "name":   user.Name,
            "mobile": user.Mobile,
            "photo":  user.Photo,
        },
    }, nil
}
