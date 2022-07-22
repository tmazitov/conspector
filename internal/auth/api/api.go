package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tmazitov/conspektor_backend.git/internal/auth/storage"
	"github.com/tmazitov/conspektor_backend.git/pkg/jwt"
)

type Api struct {
	Router  *gin.Engine
	Storage *storage.Storage
	Jwt     *jwt.JwtStorage
}

func NewApi(router *gin.Engine, storage *storage.Storage, jwt *jwt.JwtStorage) *Api {

	api := Api{Router: router, Storage: storage, Jwt: jwt}

	group := router.Group("/auth")
	group.POST("/login", api.login)
	group.POST("/create", api.create)
	return &api
}
