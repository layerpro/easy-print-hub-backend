package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/layerpro/upload-download-backend/configs"
)

func Connection(config configs.Config) (*sql.DB, error) {
	conStr := fmt.Sprintf(
		`user=%s password=%s dbname=%s port=%s sslmode=%s`,
		config.Database.User,
		config.Database.Password,
		config.Database.Name,
		config.Database.Port,
		config.Database.SllMode,
	)
	con, err := sql.Open(config.Database.Driver, conStr)
	if err != nil {
		log.Printf(`Failed to connection database: %v`, err)
		return nil, err
	}

	con.SetMaxOpenConns(config.Database.MaxOpenCon)
	con.SetMaxIdleConns(config.Database.MaxIdleCon)
	con.SetConnMaxLifetime(config.Database.MaxLifeTime)

	err = con.Ping()
	if err != nil {
		return nil, fmt.Errorf(`failed to ping database: %v`, err)
	}

	log.Printf(`Database connection success.`)
	return con, nil
}
