package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"sync"
	"time"
)

var (
	clt  *client
	once sync.Once
)

type Addresser interface {
	Address() string
	Password() string
	DB() int
}

func init() {
	_ = NewClient(defaultConfig)
}

func NewClient(cfg Addresser) *client {
	once.Do(func() {
		clt = connect(cfg.Address(), cfg.Password(), cfg.DB())
	})

	return clt
}

func Client() *client {
	return clt
}

type client struct {
	*redis.Client
}

func connect(addr, pass string, db int) *client {
	return &client{redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       db,
	})}
}

func (c *client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.Client.Set(ctx, key, value, expiration).Err()
}

func (c *client) Get(ctx context.Context, key string) (string, error) {
	return c.Client.Get(ctx, key).Result()
}
