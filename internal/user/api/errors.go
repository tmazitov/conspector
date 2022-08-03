package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrBadRequest(err error, c *gin.Context) {
	log.Println(err)
	c.JSON(http.StatusBadRequest, gin.H{
		"status": "bad request",
	})
}

func ErrInternalServer(err error, c *gin.Context) {
	log.Println(err)
	c.JSON(http.StatusInternalServerError, gin.H{
		"status": "internal server error",
	})
}

func ErrDataConflict(err error, c *gin.Context) {
	log.Println(err)
	c.JSON(http.StatusConflict, gin.H{
		"status": "conflict",
	})
}

func ErrUnauthorized(err error, c *gin.Context) {
	log.Println(err)
	c.JSON(http.StatusUnauthorized, gin.H{
		"status": "unauthorized",
	})
}
