package configs

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	APP_PORT         string
	DB_DRIVER        string
	DB_PORT          string
	DB_USER          string
	DB_PASSWORD      string
	DB_NAME          string
	DB_SLLMODE       string
	DB_MAX_OPEN_CON  int
	DB_MAX_IDLE_CON  int
	DB_MAX_LIFE_TIME time.Duration
}

var AppConfig Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(`Error load .env file. %v`, err)
	}

	maxOpenCon, err := strconv.Atoi(os.Getenv(`DB_MAX_OPEN_CON`))
	if err != nil {
		maxOpenCon = 10
	}

	maxIdleCon, err := strconv.Atoi(os.Getenv(`DB_MAX_IDLE_CON`))
	if err != nil {
		maxIdleCon = 5
	}

	maxLifeTime, err := time.ParseDuration(os.Getenv(`DB_MAX_LIFE_TIME`))
	if err != nil {
		maxLifeTime = 30 * time.Minute
	}

	AppConfig = Config{
		APP_PORT:         os.Getenv(`APP_PORT`),
		DB_DRIVER:        os.Getenv(`DB_DRIVER`),
		DB_PORT:          os.Getenv(`DB_PORT`),
		DB_USER:          os.Getenv(`DB_USER`),
		DB_PASSWORD:      os.Getenv(`DB_PASSWORD`),
		DB_NAME:          os.Getenv(`DB_NAME`),
		DB_SLLMODE:       os.Getenv(`DB_SLLMODE`),
		DB_MAX_OPEN_CON:  maxOpenCon,
		DB_MAX_IDLE_CON:  maxIdleCon,
		DB_MAX_LIFE_TIME: maxLifeTime,
	}
}
