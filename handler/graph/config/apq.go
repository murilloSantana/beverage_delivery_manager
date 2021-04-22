package config

import (
	redisCli "beverage_delivery_manager/config/redis"
	"beverage_delivery_manager/config/settings"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

const apqPrefix = "apq:"

type cacheAPQ struct {
	client redis.UniversalClient
	ttl    time.Duration
}

func NewAPQ(sts settings.RedisSettings) (*cacheAPQ, error) {
	client, err := redisCli.NewClient(sts)

	if err != nil {
		return nil, err
	}

	return &cacheAPQ{client: client, ttl: 1 * time.Minute}, nil
}

func (c *cacheAPQ) Add(ctx context.Context, key string, value interface{}) {
	c.client.Set(ctx, fmt.Sprintf("%v%v", apqPrefix, key), value, c.ttl)
}

func (c *cacheAPQ) Get(ctx context.Context, key string) (interface{}, bool) {
	s, err := c.client.Get(ctx, fmt.Sprintf("%v%v", apqPrefix, key)).Result()
	if err != nil {
		return struct{}{}, false
	}
	return s, true
}
