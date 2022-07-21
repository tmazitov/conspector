package config

import (
	"database/sql"
	"flag"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gobwas/flagutil"
	"github.com/gobwas/flagutil/parse/pargs"
)

type Config struct {
	db DB
}

func (c *Config) Setup(ctx *gin.Context) {
	flags := flag.NewFlagSet("conspektor", flag.ExitOnError)

	c.db = DB{}

	flags.StringVar(&c.db.host,
		"db_url", "127.0.0.1:5432",
		"url for conection to db",
	)

	flagutil.Parse(ctx, flags,
		flagutil.WithParser(&pargs.Parser{
			Args: os.Args[1:],
		}))
}

func (c *Config) SetupDB() (*sql.DB, error) {
	return c.db.setup()
}
