package main

import (
	"log"

	"github.com/gin-gonic/gin"
	config "github.com/tmazitov/conspektor_backend.git/config/aaa"
	aaa "github.com/tmazitov/conspektor_backend.git/internal/aaa"
)

func main() {

	ctx := &gin.Context{}
	conf := config.Config{Path: "../../config/aaa/config.json"}
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

	aaa.SetupService(r, db, redis, conf)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
