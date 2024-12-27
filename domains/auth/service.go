package auth

import (
	"database/sql"
	"errors"
	"log"

	"github.com/layerpro/upload-download-backend/utils"
)

type Service struct {
	repo Repository
	jwt  *utils.JwtConfig
}

func NewService(repo Repository, jwtConfig *utils.JwtConfig) Service {
	return Service{
		repo: repo,
		jwt:  jwtConfig,
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
