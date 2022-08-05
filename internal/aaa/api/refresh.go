package api

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmazitov/conspektor_backend.git/pkg/token"
)

type refreshRequest struct {
	Refresh string `json:"refresh" binding:"required" validate:"max=303"`
}

type refreshResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func (a *Api) refresh(c *gin.Context) {
	var (
		json         refreshRequest
		newtokenPair map[string]string
		oldTokenPair map[string]string
		err          error
	)
	if err := c.BindJSON(&json); err != nil {
		ErrBadRequest(err, c)
		return
	}
	oldTokenPair = make(map[string]string)

	oldTokenPair["refresh"] = json.Refresh
	access := token.GetAccess(c)
	if access == "" {
		ErrUnauthorized(errors.New("no access token"), c)
		return
	}

	oldTokenPair["access"] = access

	newtokenPair, err = a.Jwt.RefreshTokenPair(c, oldTokenPair)
	if err != nil {
		ErrBadRequest(err, c)
		return
	}

	c.JSON(http.StatusOK, refreshResponse{
		Refresh: newtokenPair["refresh"],
		Access:  newtokenPair["access"],
	})
}
