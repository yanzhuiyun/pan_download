package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"pandownload/settings"
)

var (
	client *redis.Client
)

func Init() (err error) {
	redisConfig := settings.Redis()
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port),
		Password: redisConfig.Password,
		DB:       redisConfig.Dbnum,
	})
	_, err = client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}
	done := make(chan struct{})
	go func() {
		DeleteFileStatusByZero(done)
	}()
	return
}
