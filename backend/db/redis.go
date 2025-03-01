package db

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/redis/go-redis/v9"
)

// RedisClientWrapper provides a global Redis client
type RedisClientWrapper struct {
	client *redis.Client
}

// Global instance of RedisClientWrapper
var (
	redisInstance *RedisClientWrapper
	once          sync.Once
)

// InitRedis initializes the Redis client and ensures it's only set once
func InitRedis() {
	once.Do(func() {
		redisInstance = &RedisClientWrapper{
			client: redis.NewClient(&redis.Options{
				Addr: "localhost:6379",
			}),
		}
		log.Println("Redis client initialized")
	})
}

// GetRedisClient provides access to the global Redis client
func GetRedisClient() *RedisClientWrapper {
	if redisInstance == nil {
		InitRedis()
	}
	return redisInstance
}

// Save stores a key-value pair in Redis
func (r *RedisClientWrapper) Save(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return r.client.Set(ctx, key, value, expiration).Err()
}

// Get retrieves a value from Redis
func (r *RedisClientWrapper) Get(ctx context.Context, key string) (string, error) {
	return r.client.Get(ctx, key).Result()
}
