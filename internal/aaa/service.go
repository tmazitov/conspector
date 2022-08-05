package aaa

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/tmazitov/conspektor_backend.git/config/aaa"
	"github.com/tmazitov/conspektor_backend.git/internal/aaa/api"
	"github.com/tmazitov/conspektor_backend.git/internal/aaa/storage"
	"github.com/tmazitov/conspektor_backend.git/internal/middleware"
	"github.com/tmazitov/conspektor_backend.git/pkg/jwt"
)

func SetupService(router *gin.Engine, dbConn *sql.DB, redis *redis.Client, conf aaa.Config) {

	storage := storage.NewStorage(dbConn)
	jwt := jwt.NewStorage(redis, conf.GetSecret())
	middleware := middleware.Middleware{JWT: jwt}

	api.SetupApi(router, storage, jwt, &middleware)
}
