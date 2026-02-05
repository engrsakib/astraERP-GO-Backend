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

type LoginResponse struct {
    User         map[string]interface{} `json:"user"`         
    AccessToken  string                 `json:"accessToken"`  
    RefreshToken string                 `json:"refreshToken"` 
}

func NewLoginService(repo *loginRepo.LoginRepository) *LoginService {
    return &LoginService{Repo: repo}
}



func (s *LoginService) Login(req dto.UserLoginRequest) (*LoginResponse, error) {

    user, err := s.Repo.FindByMobile(req.Mobile)
    if err != nil {
       
        return nil, errors.New("invalid mobile or password")
    }

    // fmt.Println("REQ MOBILE:", req.Mobile)
    // fmt.Println("DB USER MOBILE:", user.Mobile)

    if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
        return nil, errors.New("invalid mobile or password")
    }

    payload := map[string]interface{}{
        "id":     user.ID,
        "name":   user.Name,
        "mobile": user.Mobile,
        "photo":  user.Photo,
    }

    accessToken, err := utils.GenerateToken(payload, "ACCESS_TOKEN_EXP")
    if err != nil {
        return nil, err
    }

    
    refreshToken, err := utils.GenerateToken(payload, "REFRESH_TOKEN_EXP")
    if err != nil {
        return nil, err
    }

   
    return &LoginResponse{
        User: map[string]interface{}{
            "id":     user.ID,
            "name":   user.Name,
            "mobile": user.Mobile,
            "photo":  user.Photo,
        },
        AccessToken:  accessToken,
        RefreshToken: refreshToken,
    }, nil
}
