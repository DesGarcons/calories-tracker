package redis

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	MaxRetries  int
	DialTimeout time.Duration
	Timeout     time.Duration

	redisClient *redis.Client
}

const (
	_defaultMaxRetries  = 5
	_defaultDialTimeout = 10 * time.Second
	_defaultTimeout     = 5 * time.Second
)

func NewRedis(url string) (*Redis, error) {
	r := &Redis{
		MaxRetries:  _defaultMaxRetries,
		DialTimeout: _defaultDialTimeout,
		Timeout:     _defaultTimeout,
	}

	opts, err := redis.ParseURL(url)
	if err != nil {
		return nil, fmt.Errorf("redis - NewRedis - redis.ParseURL: %w", err)
	}
	opts.MaxRetries = r.MaxRetries
	opts.DialTimeout = r.DialTimeout
	opts.ReadTimeout = r.Timeout
	opts.WriteTimeout = r.Timeout

	r.redisClient = redis.NewClient(opts)

	return r, nil
}

func (r *Redis) Close() error {
	if r.redisClient != nil {
		r.redisClient.Close()
	}
	return nil
}
