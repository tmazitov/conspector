package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmazitov/conspektor_backend.git/pkg/token"
)

type logoutRequest struct {
	Refresh string `json:"refresh" binding:"required" validate:"max=303"`
}

func (a *Api) logout(c *gin.Context) {
	var (
		json      logoutRequest
		tokenPair map[string]string
		err       error
	)
	if err = c.BindJSON(&json); err != nil {
		ErrBadRequest(err, c)
		return
	}

	tokenPair = make(map[string]string)

	tokenPair["refresh"] = json.Refresh
	access := token.GetAccess(c)
	if access == "" {
		ErrUnauthorized(errors.New("no access token"), c)
		return
	}

	tokenPair["access"] = access

	if err = a.Jwt.DeleteTokenPair(c, tokenPair); err != nil {
		ErrInternalServer(err, c)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
