package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type TokenStore struct {
	client *redis.Client
	ctx    context.Context
}

func NewTokenStore(client *redis.Client) *TokenStore {
	return &TokenStore{
		client: client,
		ctx:    context.Background(),
	}
}

func (ts *TokenStore) SaveToken(userID string, token string, expiration time.Duration) error {
	return ts.client.Set(ts.ctx, userID, token, expiration).Err()
}

func (ts *TokenStore) GetToken(userID string) (string, error) {
	return ts.client.Get(ts.ctx, userID).Result()
}

func (ts *TokenStore) DeleteToken(userID string) error {
	return ts.client.Del(ts.ctx, userID).Err()
}

func (ts *TokenStore) TokenExists(userID string) (bool, error) {
	_, err := ts.client.Get(ts.ctx, userID).Result()
	if err == redis.Nil {
		return false, nil
	}
	return err == nil, err
}