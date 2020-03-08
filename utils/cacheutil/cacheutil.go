package cacheutil

import (
	"log"
	"github.com/imran103019/audit_logs/conn"
	"strings"
	"time"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

// SetCache set cache on redis
func SetCache(key string, val string, ttl ...time.Duration) error {

	var cacheTTL time.Duration

	if len(ttl) == 0 { // if optional ttl for cache is not provided, read the ttl from consul as default
		cacheTTL = viper.GetDuration("redis.cache_ttl") * time.Second
	} else {
		cacheTTL = ttl[0]
		if cacheTTL == (0 * time.Second) { //if 0 second cache is provided assign the default
			cacheTTL = viper.GetDuration("redis.cache_ttl") * time.Second
		}
	}

	rds := conn.DefaultRedis()

	key = strings.Join([]string{viper.GetString("redis.prefix"), key}, "")

	if err := rds.Set(key, val, cacheTTL).
		Err(); err != nil {
		log.Println("failed to set cache: ", err)
		return err
	}

	return nil
}

// GetCache returns a cache value against the key provided from redis
func GetCache(key string) (string, error) {
	rds := conn.DefaultRedis()

	key = strings.Join([]string{viper.GetString("redis.prefix"), key}, "")

	bb, erR := rds.Get(key).Result()
	if erR != nil {
		if erR != redis.Nil {
			log.Println("redis failed: ", erR)
			return "", erR
		}
	}

	return bb, nil

}

// ClearCache clears a cache matching the patter of the redis key
func ClearCache(pattern string) error {
	rds := conn.DefaultRedis()

	pattern = strings.Join([]string{viper.GetString("redis.prefix"), pattern}, "")

	keys, err := rds.Keys(pattern).Result()

	if err != nil {
		if err != redis.Nil {
			log.Println("redis failed: ", err)
			return err
		}
	}

	if len(keys) > 0 {
		if err := rds.Del(keys...).Err(); err != nil {
			return err
		}
	}

	return nil
}

// Exists checks if a redis key exists or not, because the function redis package provides is sh*t
func Exists(key string) (bool, error) {
	rds := conn.DefaultRedis()

	key = strings.Join([]string{viper.GetString("redis.prefix"), key}, "")

	_, err := rds.Get(key).Result()

	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// BuildKey builds a valid key for redis cache
func BuildKey(keywords ...string) string {
	kk := []string{}
	for _, kw := range keywords {
		if kw != "" {
			kk = append(kk, kw)
		}
	}

	key := strings.Join(kk, "_")

	return key
}
