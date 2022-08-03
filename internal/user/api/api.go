package api

import (
	"github.com/gin-gonic/gin"
	"github.com/tmazitov/conspektor_backend.git/internal/middleware"
	"github.com/tmazitov/conspektor_backend.git/internal/user/storage"
	"github.com/tmazitov/conspektor_backend.git/pkg/jwt"
)

type Api struct {
	Router  *gin.Engine
	middle  *middleware.Middleware
	Storage *storage.Storage
	Jwt     *jwt.Storage
}

func NewApi(router *gin.Engine, storage *storage.Storage, jwt *jwt.Storage, middle *middleware.Middleware) *Api {

	api := Api{Router: router, Storage: storage, Jwt: jwt, middle: middle}

	auth := router.Group("/auth")
	auth.POST("/login", api.login)
	auth.POST("/create", api.create)
	auth.POST("/logout", api.middle.Authorized(), api.logout)
	auth.POST("/refresh", api.refresh)

	user := router.Group("/user").Use(api.middle.Authorized())
	user.GET("/:uid", api.profile)

	return &api
}
