package mysql

import (
	"context"
	"errors"
	"github.com/jmoiron/sqlx"
	"math"
	"sync"
	"time"
)

const maxConnectAttempt = 5

var (
	pool               *connection
	once               sync.Once
	errTooManyAttempts = errors.New("too many attempts to connect")
)

type DSNer interface {
	FormatDSN() string
}

// NewConnectionPool create main struct with connections
func NewConnectionPool(cfg DSNer) *connection {
	once.Do(func() {
		pool = &connection{
			cfg:        cfg,
			readPool:   nil,
			readPoolMu: sync.Once{},
		}
	})

	return pool
}

type connection struct {
	cfg        DSNer
	readPool   *sqlx.DB
	readPoolMu sync.Once
}

func (c *connection) Ping() error {
	ctx := context.Background()

	if c.readPool != nil {
		if err := c.readPool.PingContext(ctx); err != nil {
			return err
		}
	}

	return nil
}

// ReadPool gets read connection pool
func (c *connection) ReadPool() *sqlx.DB {
	c.readPoolMu.Do(func() {
		conn, err := connect(c.cfg.FormatDSN())
		if err != nil {
			panic(err)
		}

		c.readPool = conn
	})

	return c.readPool
}

// WritePool gets write connection pool
func (c *connection) WritePool() *sqlx.DB {
	return c.ReadPool()
}

func connect(dsn string) (*sqlx.DB, error) {
	var getReadConnAttempts float64

start:
	if getReadConnAttempts > maxConnectAttempt {
		return nil, errTooManyAttempts
	}

	conn, err := sqlx.Open("mysql", dsn)
	if err != nil {
		time.Sleep(time.Duration(int(math.Pow(2, getReadConnAttempts))) * time.Second)
		getReadConnAttempts++

		goto start
	}

	return conn, nil
}
