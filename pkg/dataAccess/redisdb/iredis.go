package redisdb

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisMethods interface {
	Set(string, interface{}, time.Duration) *redis.StatusCmd
	Get(string) *redis.StringCmd
}
