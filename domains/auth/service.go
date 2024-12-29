package auth

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/layerpro/easy-print-hub-backend/utils"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	repo  Repository
	jwt   *utils.JwtConfig
	redis *redis.Client
}

func NewService(repo Repository, jwtConfig *utils.JwtConfig, redis *redis.Client) Service {
	return Service{
		repo:  repo,
		jwt:   jwtConfig,
		redis: redis,
	}
}

var ErrWrongEmailOrPassword = errors.New("wrong email or password")

func (s Service) SignIn(data SignIn) (*ResponseSignIn, error) {
	user, err := s.repo.GetUserByEmail(data.Email)
	if err == sql.ErrNoRows {
		return nil, ErrWrongEmailOrPassword
	}
	if err != nil {
		log.Printf(`Failed get user by Emil %v`, err)
		return nil, err
	}

	err = utils.CompareHashAndPassword(user.Password, data.Password)
	if err != nil {
		return nil, ErrWrongEmailOrPassword
	}

	payload := utils.PayloadJwt{
		ID:   user.ID,
		Name: user.Name,
	}

	access_token, err := s.jwt.GenerateAccessToken(payload)
	if err != nil {
		return nil, err
	}

	response := ResponseSignIn{
		AccessToken: access_token,
	}

	return &response, nil
}

func (s Service) SignOut(token string) error {
	redisKey := fmt.Sprintf(`expired-%s`, token)
	redisDuration := time.Second * time.Duration(s.jwt.GetTtl())

	err := s.redis.Set(context.Background(), redisKey, true, redisDuration).Err()
	if err != nil {
		log.Printf(`Error set redis. %v`, err)
		return err
	}
	return nil
}
