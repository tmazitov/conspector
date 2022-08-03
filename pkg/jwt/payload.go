package jwt

import (
	"time"
)

type Payload struct {
	UID       string    `json:"id"`
	Username  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(username string, uid string, duration time.Duration) *Payload {
	payload := &Payload{
		UID:       uid,
		Username:  username,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
	return payload
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrInvalidToken
	}
	return nil
}
