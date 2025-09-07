package cache

import (
	"calories-tracker/pkg/redis"
)

type BotCache struct {
	*redis.Redis
}
