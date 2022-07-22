package jwt

import (
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/golang-jwt/jwt/v4"
)

type JwtStorage struct {
	redis  *redis.Client
	secret []byte
}

func NewStorage(redis *redis.Client, secret []byte) *JwtStorage {
	return &JwtStorage{redis: redis, secret: secret}
}

func (js *JwtStorage) createToken(username string, uid string, duration time.Duration) (string, error) {
	payload := NewPayload(username, uid, duration)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString(js.secret)
}

func (js *JwtStorage) CreateTokenPair(username string, uid string) (map[string]string, error) {

	var (
		err error
	)

	tokenPair := make(map[string]string)

	if tokenPair["access"], err = js.createToken(username, uid, time.Minute*15); err != nil {
		return nil, err
	}
	if tokenPair["refresh"], err = js.createToken(username, uid, time.Hour*24*30); err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func (js *JwtStorage) Check(token string) (bool, error) {
	payload, err := js.VerifyToken(token)
	if err != nil {
		return false, err
	}

	return payload == nil, nil
}

func (js *JwtStorage) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return js.secret, nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
