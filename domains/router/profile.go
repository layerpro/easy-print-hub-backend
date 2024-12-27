package router

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/layerpro/upload-download-backend/domains/profile"
	"github.com/layerpro/upload-download-backend/utils"
)

func SetupProfileRouter(router *mux.Router, db *sql.DB, jwt *utils.JwtConfig) *mux.Router {
	profileRepo := profile.NewRepository(db)
	profileService := profile.NewService(profileRepo)
	profileHandler := profile.NewHandler(profileService)

	profileRouter := router.NewRoute().Subrouter()
	profileRouter.Use(utils.AuthMiddleware(jwt))

	profileRouter.HandleFunc(`/profile`, profileHandler.GetProfile).Methods(http.MethodGet)

	return profileRouter
}
