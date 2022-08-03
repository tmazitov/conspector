package auth

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	c "github.com/tmazitov/conspektor_backend.git/config"
	"github.com/tmazitov/conspektor_backend.git/internal/middleware"
	"github.com/tmazitov/conspektor_backend.git/internal/user/api"
	"github.com/tmazitov/conspektor_backend.git/internal/user/storage"
	"github.com/tmazitov/conspektor_backend.git/pkg/jwt"
)

type UserService struct {
	Api *api.Api
}

func NewAuthService(router *gin.Engine, dbConn *sql.DB, redis *redis.Client, conf c.Config) *UserService {

	storage := storage.NewStorage(dbConn)
	jwt := jwt.NewStorage(redis, conf.GetSecret())
	middleware := middleware.Middleware{JWT: jwt}

	api := api.NewApi(router, storage, jwt, &middleware)
	service := UserService{
		Api: api,
	}

	return &service
}
