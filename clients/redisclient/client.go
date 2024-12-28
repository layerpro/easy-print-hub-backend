package redisclient

import (
	"fmt"

	"github.com/layerpro/upload-download-backend/configs"
	"github.com/redis/go-redis/v9"
)

func Connect(config configs.Config) *redis.Client {
	rdConnStr := fmt.Sprintf(
		"redis://:%s@%s:%s",
		config.Redis.Password,
		config.Redis.Host,
		config.Redis.Port,
	)
	opt, err := redis.ParseURL(rdConnStr)
	if err != nil {
		panic(err)
	}

	rdb := redis.NewClient(opt)
	return rdb
}
