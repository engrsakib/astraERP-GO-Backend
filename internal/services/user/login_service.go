package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	dto "github.com/engrsakib/erp-system/internal/dto/user/login"
	loginRepo "github.com/engrsakib/erp-system/internal/repository/login"
	"github.com/engrsakib/erp-system/internal/utils"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
)

type LoginService struct {
	Repo *loginRepo.LoginRepository
	Rdb  *redis.Client
}

type LoginResponse struct {
	User         map[string]interface{} `json:"user"`
	AccessToken  string                 `json:"accessToken"`
	RefreshToken string                 `json:"refreshToken"`
}

func NewLoginService(repo *loginRepo.LoginRepository, rdb *redis.Client) *LoginService {
	return &LoginService{
		Repo: repo,
		Rdb:  rdb,
	}
}

func (s *LoginService) Login(req dto.UserLoginRequest) (*LoginResponse, error) {
	ctx := context.Background()
	lockKey := "login_attempts:" + req.Mobile

	attempts, _ := s.Rdb.Get(ctx, lockKey).Int()
	if attempts >= 5 {
		ttl, _ := s.Rdb.TTL(ctx, lockKey).Result()
		return nil, fmt.Errorf("account locked due to too many failed attempts. try again in %v", ttl)
	}

	user, err := s.Repo.FindByMobile(req.Mobile)
	if err != nil {
		return nil, errors.New("invalid mobile or password")
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		newAttempts, _ := s.Rdb.Incr(ctx, lockKey).Result()
		if newAttempts == 1 {
			s.Rdb.Expire(ctx, lockKey, 45*time.Minute)
		}
		if newAttempts >= 5 {
			s.Rdb.Expire(ctx, lockKey, 45*time.Minute)
		}
		return nil, errors.New("invalid mobile or password")
	}

	s.Rdb.Del(ctx, lockKey)

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