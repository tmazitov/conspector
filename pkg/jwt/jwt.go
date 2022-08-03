package jwt

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Storage struct {
	redis  *redis.Client
	secret []byte
}

func NewStorage(redis *redis.Client, secret []byte) *Storage {
	return &Storage{redis: redis, secret: secret}
}

func (s *Storage) appendToActive(c *gin.Context, tokenPair map[string]string) {
	s.redis.LPush(c, "active:access", tokenPair["access"])   // expire not working
	s.redis.LPush(c, "active:refresh", tokenPair["refresh"]) // expire not working
}

func (s *Storage) removeInActive(c *gin.Context, tokenPair map[string]string) {
	s.redis.LRem(c, "active:access", 1, tokenPair["access"])
	s.redis.LRem(c, "active:refresh", 1, tokenPair["refresh"])
}

func (s *Storage) isExists(c *gin.Context, listName string, token string) error {
	err := s.redis.LPos(c, listName, token, redis.LPosArgs{}).Err()
	fmt.Println("check ", listName)
	if err != nil {
		return ErrInvalidToken
	}

	return nil
}

func (s *Storage) DeleteTokenPair(c *gin.Context, tokenPair map[string]string) error {
	s.removeInActive(c, tokenPair)
	return nil
}

// ok    : refresh -> write new pair to redis
// to do : create  -> write new pair to redis
// to do : delete  -> delete new pair from redis
