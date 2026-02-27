package idempotency

import (
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/net/context"
)

type IdempotencyKeyManager struct {
	redisClient *redis.Client
	expiration  time.Duration
}

func NewIdempotencyKeyManager(redisClient *redis.Client, expiration time.Duration) *IdempotencyKeyManager {
	return &IdempotencyKeyManager{
		redisClient: redisClient,
		expiration:  expiration,
	}
}

func (ikm *IdempotencyKeyManager) Set(key string, value string) error {
	ctx := context.Background()
	err := ikm.redisClient.Set(ctx, key, value, ikm.expiration).Err()
	if err != nil {
		return errors.New("failed to set idempotency key: " + err.Error())
	}
	return nil
}

func (ikm *IdempotencyKeyManager) Get(key string) (string, error) {
	ctx := context.Background()
	value, err := ikm.redisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		return "", errors.New("failed to get idempotency key: " + err.Error())
	}
	return value, nil
}