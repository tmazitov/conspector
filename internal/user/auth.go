package auth

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	c "github.com/tmazitov/conspektor_backend.git/config"
	"github.com/tmazitov/conspektor_backend.git/internal/user/api"
	"github.com/tmazitov/conspektor_backend.git/internal/user/storage"
	"github.com/tmazitov/conspektor_backend.git/pkg/jwt"
)

type AuthService struct {
	Api *api.Api
}

func NewAuthService(router *gin.Engine, dbConn *sql.DB, redis *redis.Client, conf c.Config) *AuthService {

	storage := storage.NewStorage(dbConn)
	jwt := jwt.NewStorage(redis, conf.GetSecret())

	api := api.NewApi(router, storage, jwt)
	service := AuthService{
		Api: api,
	}

	return &service
}
