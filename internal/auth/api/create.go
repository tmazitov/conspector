package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmazitov/conspektor_backend.git/internal/auth/dto"
	"github.com/tmazitov/conspektor_backend.git/pkg/hash"
)

type Request struct {
	Username string `json:"username" binding:"required" validate:"min=3,max=20"`
	Password string `json:"password" binding:"required" validate:"min=8,max=20"`
	Email    string `json:"email"    binding:"required" validate:"email"`
}

func (a *Api) create(c *gin.Context) {

	var (
		json Request
		err  error
	)
	if err := c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": c.Error(err),
		})
		return
	}

	if err = a.Storage.CheckEmail(json.Email); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": c.Error(err),
		})
		return
	}

	if a.Storage.CheckUsername(json.Username); err != nil {
		c.JSON(http.StatusConflict, gin.H{
			"status": c.Error(err),
		})
		return
	}

	upload := dto.CreateUser{
		Username: json.Username,
		Password: hash.GenerateSha256(json.Password),
		Email:    json.Email,
	}

	if err := a.Storage.UserCreate(upload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": c.Error(err),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user was created",
	})
}
