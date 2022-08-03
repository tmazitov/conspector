package jwt

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (s *Storage) createToken(username string, uid string, duration time.Duration) (string, error) {
	payload := NewPayload(username, uid, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString(s.secret)
}

func (s *Storage) CreateTokenPair(c *gin.Context, username string, uid string) (map[string]string, error) {

	var (
		err error
	)

	tokenPair := make(map[string]string)

	if tokenPair["access"], err = s.createToken(username, uid, time.Minute*15); err != nil {
		return nil, err
	}
	if tokenPair["refresh"], err = s.createToken(username, uid, time.Hour*24*30); err != nil {
		return nil, err
	}

	s.appendToActive(c, tokenPair)

	return tokenPair, nil
}
