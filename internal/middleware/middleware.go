package middleware

import "github.com/tmazitov/conspektor_backend.git/pkg/jwt"

type Middleware struct {
	JWT *jwt.Storage
}
