package router

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/layerpro/upload-download-backend/domains/auth"
	"github.com/layerpro/upload-download-backend/utils"
)

func SetupAuthrouter(router *mux.Router, db *sql.DB, jwt *utils.JwtConfig) *mux.Router {
	authRepo := auth.NewRepository(db)
	authService := auth.NewService(authRepo, jwt)
	authHandler := auth.NewHandler(authService)

	router.HandleFunc(`/sign-in`, authHandler.SignIn).Methods(http.MethodPost)

	return router
}
