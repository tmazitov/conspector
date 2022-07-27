package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmazitov/conspektor_backend.git/internal/auth/dto"
	"github.com/tmazitov/conspektor_backend.git/internal/auth/models"
	"github.com/tmazitov/conspektor_backend.git/pkg/hash"
)

type loginRequest struct {
	Username string `json:"username" binding:"required" validate:"min=3,max=20"`
	Password string `json:"password" binding:"required" validate:"min=8,max=20"`
}

type loginResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func (a *Api) login(c *gin.Context) {
	var (
		json      loginRequest
		user      models.User
		tokenPair map[string]string
		err       error
	)
	if err = c.BindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": c.Error(err),
		})
		return
	}

	json.Password = hash.GenerateSha256(json.Password)
	dto := dto.CheckPassword{
		Username: json.Username,
		Password: json.Password,
	}

	if err = a.Storage.User.CheckPassword(dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": c.Error(err),
		})
		return
	}

	if user, err = a.Storage.User.Profile(json.Username); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": c.Error(err),
		})
		return
	}

	if tokenPair, err = a.Jwt.CreateTokenPair(user.Username, user.UID); err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status": c.Error(err),
		})
		return
	}

	c.JSON(http.StatusOK, loginResponse{Access: tokenPair["access"], Refresh: tokenPair["refresh"]})
}
