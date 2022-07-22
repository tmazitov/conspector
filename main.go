package main

import (
	"log"

	"github.com/gin-gonic/gin"
	c "github.com/tmazitov/conspektor_backend.git/config"
	"github.com/tmazitov/conspektor_backend.git/internal/auth"
)

func main() {

	ctx := &gin.Context{}
	conf := c.Config{}
	if err := conf.Setup(ctx); err != nil {
		log.Fatal(err)
	}

	db, err := conf.SetupDB()
	if err != nil {
		log.Fatal(err)
	}

	redis := conf.SetupRedis(ctx)

	defer db.Close()
	defer redis.Close()

	r := gin.Default()

	_ = auth.NewAuthService(r, db, redis, conf)

	r.Run("localhost:8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
