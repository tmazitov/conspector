package main

import (
	"log"

	"github.com/gin-gonic/gin"
	c "github.com/tmazitov/conspektor_backend.git/config"
	"github.com/tmazitov/conspektor_backend.git/internal/auth"
)

var conf c.Config = c.Config{}

func init() {
	conf.Setup(&gin.Context{})
}

func main() {
	conn, err := conf.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	r := gin.Default()

	_ = auth.NewAuthService(r, conn)

	r.Run("localhost:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
