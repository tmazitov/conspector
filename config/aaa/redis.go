package aaa

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type RedisConfig struct {
	host   string
	pass   string
	secret string
	db     int
}

func (r *RedisConfig) setup(ctx *gin.Context) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     r.host,
		Password: r.pass, // no password set
		DB:       r.db,   // use default DB
	})

	return rdb
}
