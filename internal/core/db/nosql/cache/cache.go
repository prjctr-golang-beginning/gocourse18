package cache

import (
	"context"
	"gocourse18/internal/core/db/nosql/redis"
	"time"
)

func Store(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return redis.Client().Set(ctx, key, value, expiration)
}

func Load(ctx context.Context, key string) (string, error) {
	return redis.Client().Get(ctx, key)
}
