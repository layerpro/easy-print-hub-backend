package router

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/layerpro/upload-download-backend/domains/auth"
	"github.com/layerpro/upload-download-backend/utils"
	"github.com/redis/go-redis/v9"
)

func SetupAuthrouter(router *mux.Router, db *sql.DB, jwt *utils.JwtConfig, redis *redis.Client, authMiddleware func(http.Handler) http.Handler) *mux.Router {
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo, jwt, redis)
	authHandler := auth.NewHandler(authService)

	router.HandleFunc(`/sign-in`, authHandler.SignIn).Methods(http.MethodPost)
	router.Handle(`/sign-out`, authMiddleware(http.HandlerFunc(authHandler.SignOut))).Methods(http.MethodPost)

	return router
}
