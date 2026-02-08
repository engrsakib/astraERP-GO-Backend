package user

import (
	"errors"
	"fmt"
	"os"
	dto "github.com/engrsakib/erp-system/internal/dto/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/engrsakib/erp-system/internal/utils"
)


func (s *LoginService) RefreshToken(req dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {

	
	token, err := jwt.Parse(req.RefreshToken, func(token *jwt.Token) (interface{}, error) {
		
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired refresh token")
	}

	
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	
	userIDFloat, ok := claims["id"].(float64)
	if !ok {
		return nil, errors.New("invalid user id in token")
	}
	userID := int64(userIDFloat)


	user, err := s.Repo.FindByID(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	
	payload := map[string]interface{}{
		"id":        user.ID,
		"name":      user.Name,
		"mobile":    user.Mobile,
		"photo":     user.Photo,
		"user_type": user.UserType, 
	}

	
	newAccessToken, err := utils.GenerateToken(payload, "ACCESS_TOKEN_EXP")
	if err != nil {
		return nil, err
	}

	
	newRefreshToken, err := utils.GenerateToken(payload, "REFRESH_TOKEN_EXP")
	if err != nil {
		return nil, err
	}

	return &dto.RefreshTokenResponse{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}