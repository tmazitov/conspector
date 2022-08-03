package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmazitov/conspektor_backend.git/internal/user/models"
)

type Request struct {
	Uid string `uri:"uid" binding:"required,uuid"`
}

type Responce struct {
}

func (a *Api) profile(c *gin.Context) {

	var (
		json *Request
		user models.User
		err  error
	)

	if err = c.ShouldBindUri(&json); err != nil {
		ErrBadRequest(err, c)
		return
	}

	user, err = a.Storage.User.ProfileById(json.Uid)
	if err != nil {
		ErrInternalServer(err, c)
		return
	}

	c.JSON(http.StatusOK, user)
}
