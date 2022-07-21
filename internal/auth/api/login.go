package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Api) login(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
