package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/layerpro/easy-print-hub-backend/clients/database"
	"github.com/layerpro/easy-print-hub-backend/clients/redisclient"
	"github.com/layerpro/easy-print-hub-backend/clients/validator"
	"github.com/layerpro/easy-print-hub-backend/configs"
	"github.com/layerpro/easy-print-hub-backend/domains/router"
	"github.com/layerpro/easy-print-hub-backend/utils"
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
	redis := redisclient.Connect(conf)
	authMiddleware := utils.AuthMiddleware(jwt, redis)

	routes := mux.NewRouter()

	router.SetupAuthrouter(routes, db, jwt, redis, authMiddleware)
	router.SetupProfileRouter(routes, db, authMiddleware)

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
