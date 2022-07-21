package auth

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/tmazitov/conspektor_backend.git/internal/auth/api"
	"github.com/tmazitov/conspektor_backend.git/internal/auth/storage"
)

type AuthService struct {
	Api *api.Api
}

func NewAuthService(router *gin.Engine, conn *sql.DB) *AuthService {

	storage := storage.NewStorage(conn)
	api := api.NewApi(router, storage)
	service := AuthService{
		Api: api,
	}

	return &service
}
