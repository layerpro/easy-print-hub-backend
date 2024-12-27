package configs

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	App      App
	Database Database
	Jwt      Jwt
	Redis    Redis
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf(`Error load .env file. %v`, err)
	}
	return Config{
		App:      app(),
		Database: database(),
		Jwt:      jwt(),
		Redis:    redis(),
	}
}

type App struct {
	Port string
}

func app() App {
	port := os.Getenv(`APP_PORT`)
	if port == `` {
		port = `3000`
	}

	return App{
		Port: port,
	}
}

type Database struct {
	Driver      string
	Port        string
	User        string
	Password    string
	Name        string
	SllMode     string
	MaxOpenCon  int
	MaxIdleCon  int
	MaxLifeTime time.Duration
}

func database() Database {
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

	return Database{
		Driver:      os.Getenv(`DB_DRIVER`),
		Port:        os.Getenv(`DB_PORT`),
		User:        os.Getenv(`DB_USER`),
		Password:    os.Getenv(`DB_PASSWORD`),
		Name:        os.Getenv(`DB_NAME`),
		SllMode:     os.Getenv(`DB_SLLMODE`),
		MaxOpenCon:  maxOpenCon,
		MaxIdleCon:  maxIdleCon,
		MaxLifeTime: maxLifeTime,
	}
}

type Jwt struct {
	SecretAccessToken string
	TtlAccessToken    int
}

func jwt() Jwt {
	ttlAccessToken, err := strconv.Atoi(os.Getenv(`JWT_TTL_ACCESS_TOKEN`))
	if err != nil {
		ttlAccessToken = 7200
	}

	return Jwt{
		SecretAccessToken: os.Getenv(`JWT_SECRET_ACCESS_TOKEN`),
		TtlAccessToken:    ttlAccessToken,
	}
}

type Redis struct {
	Host     string
	Port     string
	Password string
}

func redis() Redis {
	return Redis{
		Host:     os.Getenv(`REDIS_HOST`),
		Port:     os.Getenv(`REDIS_PORT`),
		Password: os.Getenv(`REDIS_PASSWORD`),
	}
}
