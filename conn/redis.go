package conn

import (
	"github.com/imran103019/audit_logs/config"

	"github.com/go-redis/redis"
)

// RedisClient holds the redis client instance
type RedisClient struct {
	*redis.Client
}

// Redis is an instance *redis.Client{}
var redisCl RedisClient

// Setup assigns redis.Client interface based on config to RedisClient
func (r *RedisClient) Setup(cfg *config.RedisCfg) {
	c := redis.NewClient(&redis.Options{
		Addr:     cfg.Address,
		Password: cfg.Password,
		DB:       cfg.DB,
	})
	r.Client = c
}

// ConnectRedis provides a connector to redis based on configurations set
func ConnectRedis() error {
	cfg := config.Redis()
	redisCl.Setup(&cfg)
	return nil
}

// DefaultRedis returns the default RedisClient currently in Use
func DefaultRedis() RedisClient {
	return redisCl
}