package rediscache

import (
	"context"
	"errors"
	"time"

	"github.com/go-hypercube/go-hypercube/cache"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func New(client *redis.Client) *RedisCache {
	return &RedisCache{client: client}
}

func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
	value, err := c.client.Get(ctx, key).Result()
	if errors.Is(err, redis.Nil) {
		return "", cache.ErrNotFound
	}
	return value, err
}

func (c *RedisCache) Set(ctx context.Context, key string, value string, ttl time.Duration) error {
	if ttl < 0 {
		return cache.ErrInvalidTTL
	}
	return c.client.Set(ctx, key, value, ttl).Err()
}

func (c *RedisCache) Delete(ctx context.Context, key string) error {
	return c.client.Del(ctx, key).Err()
}

func (c *RedisCache) Has(ctx context.Context, key string) (bool, error) {
	n, err := c.client.Exists(ctx, key).Result()
	return n > 0, err
}

func (c *RedisCache) Increment(ctx context.Context, key string, delta int64) (int64, error) {
	return c.client.IncrBy(ctx, key, delta).Result()
}

func (c *RedisCache) Expire(ctx context.Context, key string, ttl time.Duration) error {
	if ttl < 0 {
		return cache.ErrInvalidTTL
	}
	return c.client.Expire(ctx, key, ttl).Err()
}
