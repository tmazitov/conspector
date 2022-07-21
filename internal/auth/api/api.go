package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tmazitov/conspektor_backend.git/internal/auth/storage"
)

type Api struct {
	Router  *gin.Engine
	Storage *storage.Storage
}

func NewApi(router *gin.Engine, storage *storage.Storage) *Api {

	api := Api{Router: router, Storage: storage}

	group := router.Group("/auth")
	group.POST("/login", api.login)
	group.POST("/create", api.create)
	return &api
}
