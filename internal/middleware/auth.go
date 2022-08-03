package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tmazitov/conspektor_backend.git/pkg/token"
)

func (m *Middleware) Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {

		access := token.GetAccess(c)
		_, err := m.JWT.CheckAccess(c, access)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"status": "unauthorized",
			})
			return
		}
		c.Next()
	}
}
