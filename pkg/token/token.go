package token

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAccess(c *gin.Context) string {
	var accessToken string
	authetication := c.Request.Header.Get("Authorization")
	if authetication == "" {
		return ""
	}

	if strings.Contains(authetication, "Bearer") {
		accessToken = strings.Split(authetication, " ")[1]
	} else {
		accessToken = authetication
	}

	return accessToken
}
