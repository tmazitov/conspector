package jwt

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func (s *Storage) CheckAccess(c *gin.Context, token string) (bool, error) {
	return s.check(c, "active:access", token)
}

func (s *Storage) CheckRefresh(c *gin.Context, token string) (bool, error) {
	return s.check(c, "active:refresh", token)
}

func (s *Storage) check(c *gin.Context, list string, token string) (bool, error) {
	payload, err := s.verifyToken(c, list, token)
	if err != nil {
		return false, err
	}

	return payload == nil, nil
}

func (s *Storage) verifyToken(c *gin.Context, list string, token string) (*Payload, error) {
	if err := s.isExists(c, list, token); err != nil {
		return nil, ErrInvalidToken
	}

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return s.secret, nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrInvalidToken) {
			return nil, ErrInvalidToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
