package router

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/layerpro/upload-download-backend/domains/profile"
)

func SetupProfileRouter(router *mux.Router, db *sql.DB, authMiddleware func(http.Handler) http.Handler) *mux.Router {
	profileRepo := profile.NewRepository(db)
	profileService := profile.NewService(profileRepo)
	profileHandler := profile.NewHandler(profileService)

	profileRouter := router.NewRoute().Subrouter()
	profileRouter.Use(authMiddleware)

	profileRouter.HandleFunc(`/profile`, profileHandler.GetProfile).Methods(http.MethodGet)

	return profileRouter
}
