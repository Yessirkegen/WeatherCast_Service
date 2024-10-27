package cache

import (
	"context"
	"fmt"
	"time"
	"weather-service/pkg/config"

	"github.com/redis/go-redis/v9"
)

type Cache struct {
	client *redis.Client
	ctx    context.Context
}

func NewCache() *Cache {
	cfg := config.GetConfig()
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
	})
	return &Cache{client: client, ctx: context.Background()}
}

func (c *Cache) Set(key string, value string, ttl int) error {
	return c.client.Set(c.ctx, key, value, time.Duration(ttl)*time.Second).Err()
}

func (c *Cache) Get(key string) (string, error) {
	return c.client.Get(c.ctx, key).Result()
}
