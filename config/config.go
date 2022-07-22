package config

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	db    DBConfig
	redis RedisConfig
}

func (c *Config) Setup(ctx *gin.Context) error {
	jsonFile, err := os.Open("./config/config.json")
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	c.db = DBConfig{}
	c.redis = RedisConfig{}

	c.db.host = fmt.Sprint(result["database_url"])
	c.redis.secret = fmt.Sprint(result["secret_verify_key"])

	return nil
}

func (c *Config) GetSecret() []byte {
	return []byte(c.redis.secret)
}

func (c *Config) SetupDB() (*sql.DB, error) {
	return c.db.setup()
}

func (c *Config) SetupRedis(ctx *gin.Context) *redis.Client {
	return c.redis.setup(ctx)
}
