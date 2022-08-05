package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tmazitov/conspektor_backend.git/internal/aaa/storage"
	"github.com/tmazitov/conspektor_backend.git/internal/middleware"
	"github.com/tmazitov/conspektor_backend.git/pkg/jwt"
)

type Api struct {
	Router  *gin.Engine
	middle  *middleware.Middleware
	Storage *storage.Storage
	Jwt     *jwt.Storage
}

func SetupApi(router *gin.Engine, storage *storage.Storage, jwt *jwt.Storage, middle *middleware.Middleware) {

	api := Api{Router: router, Storage: storage, Jwt: jwt, middle: middle}

	auth := router.Group("/auth")
	auth.POST("/login", api.login)
	auth.POST("/create", api.create)
	auth.POST("/logout", api.middle.Authorized(), api.logout)
	auth.POST("/refresh", api.refresh)

	user := router.Group("/user").Use(api.middle.Authorized())
	user.GET("/:uid", api.profile)

	log.Println("aaa : api success")
}
