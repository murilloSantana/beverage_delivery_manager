package redis

import (
	"beverage_delivery_manager/config/settings"
	"context"
	"github.com/go-redis/redis/v8"
)

func NewClient(sts settings.RedisSettings) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:         sts.URL,
		Password:     sts.Password,
		MinIdleConns: sts.MinIdleConns,
		IdleTimeout:  sts.IdleTimeout,
		PoolSize:     sts.PoolSize,
		DB:           sts.Database,
	})

	err := client.Ping(context.Background()).Err()

	if err != nil {
		return nil, err
	}

	return client, nil
}
