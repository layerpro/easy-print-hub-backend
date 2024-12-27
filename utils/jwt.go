package utils

import (
	"errors"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/layerpro/upload-download-backend/configs"
)

type PayloadJwt struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type JwtDecodeInterface struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type JwtConfig struct {
	secretKey string
	ttl       int
}

func NewJwt(config configs.Config) *JwtConfig {
	return &JwtConfig{
		secretKey: config.Jwt.SecretAccessToken,
		ttl:       config.Jwt.TtlAccessToken,
	}
}

func (config JwtConfig) GenerateAccessToken(payload PayloadJwt) (string, error) {
	secreatKey := config.secretKey
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   payload.ID,
		"name": payload.Name,
		"exp":  time.Now().Add(time.Second * time.Duration(config.ttl)).Unix(),
		"iat":  time.Now().Unix(),
	})

	tokenString, err := token.SignedString([]byte(secreatKey))
	if err != nil {
		log.Printf("Failed genrerate access token. %v", err)
		return "", err
	}

	return tokenString, nil
}

func (config JwtConfig) VerifyAccessToken(tokenString string) (*JwtDecodeInterface, error) {
	secreatKey := config.secretKey

	claims := &JwtDecodeInterface{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, ErrUnexpectedSigningMethod
		}
		return []byte(secreatKey), nil
	})

	if errors.Is(err, jwt.ErrTokenExpired) {
		return nil, jwt.ErrTokenExpired
	}

	if err != nil {
		log.Printf("Error verifying token: %v", err)
		return nil, err
	}

	return claims, nil
}
