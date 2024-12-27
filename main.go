package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/layerpro/upload-download-backend/clients/database"
	"github.com/layerpro/upload-download-backend/clients/validator"
	"github.com/layerpro/upload-download-backend/configs"
	"github.com/layerpro/upload-download-backend/domains/router"
	"github.com/layerpro/upload-download-backend/utils"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

func main() {
	validator.InitValidator()
	conf := configs.LoadConfig()
	port := conf.App.Port

	db, err := database.Connection(conf)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	jwt := utils.NewJwt(conf)

	routes := mux.NewRouter()

	router.SetupAuthrouter(routes, db, jwt)
	router.SetupProfileRouter(routes, db, jwt)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Ubah sesuai domain front-end Anda
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := corsMiddleware.Handler(routes)

	log.Printf("Server is running on port %s", port)
	http.ListenAndServe(`:`+port, handler)
}
