package utils

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
)

var ErrUnexpectedSigningMethod = errors.New("unexpected signing method")

type contextKey string

var UserKey contextKey = "user"

type JwtService interface {
	VerifyAccessToken(tokenString string) (*JwtDecodeInterface, error)
}

func AuthMiddleware(jwtService JwtService, redis *redis.Client) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				token := GetBearerToken(r)
				if token == "" {
					ResponseError(w, http.StatusUnauthorized, ErrUnauthorized)
					return
				}

				// Gunakan layanan JWT untuk verifikasi token
				decode, err := jwtService.VerifyAccessToken(token)
				if err == jwt.ErrTokenExpired {
					ResponseError(w, http.StatusUnauthorized, ErrTokenExpired)
					return
				}
				if err != nil {
					ResponseError(w, http.StatusUnauthorized, ErrUnauthorized)
					return
				}

				checkExpired := checkTokenExpiredOnRedis(*redis, token)
				if checkExpired {
					log.Print("Token already logout")
					ResponseError(w, http.StatusUnauthorized, ErrUnauthorized)
					return
				}

				// Tambahkan data token ke context
				ctx := context.WithValue(r.Context(), UserKey, decode)

				next.ServeHTTP(w, r.WithContext(ctx))
			})
	}
}

func GetBearerToken(r *http.Request) string {
	authorization := r.Header.Get("Authorization")
	if authorization == "" || !strings.Contains(authorization, "Bearer") {
		return ""
	}
	token := strings.Split(authorization, " ")[1]
	return token
}

func checkTokenExpiredOnRedis(redis redis.Client, token string) bool {
	checkExpired := redis.Get(context.Background(), "expired-"+token).Val()
	return checkExpired == "1"
}

func UserFromContext(ctx context.Context) (*JwtDecodeInterface, error) {
	user, ok := ctx.Value(UserKey).(*JwtDecodeInterface)
	if !ok {
		log.Printf(`Failed get data user context %v`, ctx.Value(UserKey))
		return nil, fmt.Errorf("invalid user context")
	}
	return user, nil
}
