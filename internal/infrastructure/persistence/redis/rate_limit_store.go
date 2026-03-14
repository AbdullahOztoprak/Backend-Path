package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RateLimitStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewRateLimitStore(redisClient *redis.Client) *RateLimitStore {
	return &RateLimitStore{
		client: redisClient,
		ctx:    context.Background(),
	}
}

func (r *RateLimitStore) SetRateLimit(key string, limit int, duration time.Duration) error {
	_, err := r.client.Set(r.ctx, key, limit, duration).Result()
	return err
}

func (r *RateLimitStore) GetRateLimit(key string) (int, error) {
	val, err := r.client.Get(r.ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return 0, nil
		}
		return 0, err
	}

	var limit int
	_, err = fmt.Sscanf(val, "%d", &limit)
	return limit, err
}

func (r *RateLimitStore) IncrementRateLimit(key string) (int, error) {
	limit, err := r.client.Incr(r.ctx, key).Result()
	return int(limit), err
}

func (r *RateLimitStore) ResetRateLimit(key string) error {
	_, err := r.client.Del(r.ctx, key).Result()
	return err
}