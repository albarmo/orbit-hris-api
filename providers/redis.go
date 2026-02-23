package providers

import (
	"context"
	"os"

	"github.com/Caknoooo/go-gin-clean-starter/pkg/constants"
	"github.com/redis/go-redis/v9"
	"github.com/samber/do"
)

func ProvideRedis(injector *do.Injector) {
	do.ProvideNamed(injector, constants.REDISClient, func(i *do.Injector) (*redis.Client, error) {
		addr := os.Getenv("REDIS_HOST")
		if addr == "" {
			addr = "127.0.0.1:6379"
		}

		password := os.Getenv("REDIS_PASSWORD")

		opt := &redis.Options{
			Addr:     addr,
			Password: password,
			DB:       0,
		}

		client := redis.NewClient(opt)

		// quick ping to verify
		if err := client.Ping(context.Background()).Err(); err != nil {
			return nil, err
		}

		return client, nil
	})
}
