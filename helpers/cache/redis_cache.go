package cache

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// RedisCache wraps the Redis client
type RedisCache struct {
	client *redis.Client
	ctx    context.Context
}

// NewRedisCache creates a new instance of RedisCache
func NewRedisCache(addr string, password string, db int) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,     // Redis address
		Password: password, // no password set
		DB:       db,       // use default DB
	})

	return &RedisCache{
		client: rdb,
		ctx:    context.Background(),
	}
}

// Set adds an item to the cache, with an expiration time.
func (cache *RedisCache) Set(key string, value interface{}, expiration time.Duration) error {
	return cache.client.Set(cache.ctx, key, value, expiration).Err()
}

// Get retrieves an item from the cache. If the item is not found, returns nil.
func (cache *RedisCache) Get(key string) (string, error) {
	val, err := cache.client.Get(cache.ctx, key).Result()
	if err == redis.Nil {
		return "", nil // Item not found
	}
	return val, err
}

// Delete removes an item from the cache.
func (cache *RedisCache) Delete(key string) error {
	return cache.client.Del(cache.ctx, key).Err()
}
