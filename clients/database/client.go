package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/layerpro/upload-download-backend/configs"
)

func Connection() (*sql.DB, error) {
	conStr := fmt.Sprintf(
		`user=%s password=%s dbname=%s port=%s sslmode=%s`,
		configs.AppConfig.DB_USER,
		configs.AppConfig.DB_PASSWORD,
		configs.AppConfig.DB_NAME,
		configs.AppConfig.DB_PORT,
		configs.AppConfig.DB_SLLMODE,
	)
	con, err := sql.Open("postgres", conStr)
	if err != nil {
		log.Printf(`Failed to connection database: %v`, err)
		return nil, err
	}

	con.SetMaxOpenConns(configs.AppConfig.DB_MAX_OPEN_CON)
	con.SetMaxIdleConns(configs.AppConfig.DB_MAX_IDLE_CON)
	con.SetConnMaxLifetime(configs.AppConfig.DB_MAX_LIFE_TIME)

	err = con.Ping()
	if err != nil {
		return nil, fmt.Errorf(`failed to ping database: %v`, err)
	}

	log.Printf(`Database connection success.`)
	return con, nil
}
