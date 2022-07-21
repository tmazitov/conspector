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

	flagutil.Subset(flags, "db", func(sub *flag.FlagSet) {
		sub.StringVar(&c.db.username,
			"username", "postgres",
			"username for conection to db",
		)

		sub.StringVar(&c.db.password,
			"password", "postgres",
			"password for conection to db",
		)

		sub.StringVar(&c.db.name,
			"name", "postgres",
			"name for conection to db",
		)

		sub.StringVar(&c.db.host,
			"host", "127.0.0.1:5432",
			"host for conection to db",
		)
	})

	flagutil.Parse(ctx, flags,
		flagutil.WithParser(&pargs.Parser{
			Args: os.Args[1:],
		}))
}

func (c *Config) SetupDB() (*sql.DB, error) {
	return c.db.setup()
}
