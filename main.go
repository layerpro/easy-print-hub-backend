package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/layerpro/upload-download-backend/clients/database"
	"github.com/layerpro/upload-download-backend/configs"
	_ "github.com/lib/pq"
)

func main() {
	configs.LoadConfig()

	db, err := database.Connection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	fmt.Println("Hello, World!")
	port := configs.AppConfig.APP_PORT
	if port == `` {
		port = `3000`
	}
	http.ListenAndServe(`:`+port, nil)
}
