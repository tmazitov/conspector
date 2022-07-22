package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DBConfig struct {
	host string
}

func (db *DBConfig) setup() (*sql.DB, error) {

	con, err := sql.Open("postgres", db.host)
	if err != nil {
		return nil, err
	}

	return con, nil
}
