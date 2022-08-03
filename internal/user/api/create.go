package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	userDto "github.com/tmazitov/conspektor_backend.git/internal/user/dto/user"
	"github.com/tmazitov/conspektor_backend.git/pkg/hash"
)

type createRequest struct {
	Username string `json:"username" binding:"required" validate:"min=3,max=20"`
	Password string `json:"password" binding:"required" validate:"min=8,max=20"`
	Email    string `json:"email"    binding:"required" validate:"email"`
}

type createResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func (a *Api) create(c *gin.Context) {

	var (
		json      createRequest
		tokenPair map[string]string
		err       error
	)
	if err := c.BindJSON(&json); err != nil {
		ErrBadRequest(err, c)
		return
	}

	if err = a.Storage.User.CheckEmail(json.Email); err != nil {
		ErrDataConflict(err, c)
		return
	}

	if err = a.Storage.User.CheckUsername(json.Username); err != nil {
		ErrDataConflict(err, c)
		return
	}

	dto := userDto.CreateUser{
		Username: json.Username,
		UID:      uuid.New().String(),
		Password: hash.GenerateSha256(json.Password),
		Email:    json.Email,
	}

	if err = a.Storage.User.Create(dto); err != nil {
		ErrInternalServer(err, c)
		return
	}

	tokenPair, err = a.Jwt.CreateTokenPair(c, dto.Username, dto.UID)
	if err != nil {
		ErrInternalServer(err, c)
		return
	}

	c.JSON(http.StatusCreated, createResponse{
		Access:  tokenPair["access"],
		Refresh: tokenPair["refresh"],
	})
}
