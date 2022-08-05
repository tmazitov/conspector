package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	userDto "github.com/tmazitov/conspektor_backend.git/internal/aaa/dto/user"
	"github.com/tmazitov/conspektor_backend.git/internal/aaa/models"
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
		ErrBadRequest(err, c)
		return
	}

	json.Password = hash.GenerateSha256(json.Password)
	dto := userDto.CheckPassword{
		Username: json.Username,
		Password: json.Password,
	}

	if err = a.Storage.User.CheckPassword(dto); err != nil {
		ErrBadRequest(err, c)
		return
	}

	if user, err = a.Storage.User.ProfileByUsername(json.Username); err != nil {
		ErrInternalServer(err, c)
		return
	}

	if tokenPair, err = a.Jwt.CreateTokenPair(c, user.Username, user.UID); err != nil {
		ErrInternalServer(err, c)
		return
	}

	c.JSON(http.StatusOK, loginResponse{
		Access:  tokenPair["access"],
		Refresh: tokenPair["refresh"],
	})
}
