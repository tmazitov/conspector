package config

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type DB struct {
	host string
}

func (db *DB) setup() (*sql.DB, error) {

	// connStr := fmt.Sprintf("user=%v password=%v dbname=%v host=%v", db.username, db.password, db.name, db.host)
	con, err := sql.Open("postgres", db.host)
	if err != nil {
		return nil, err
	}

	return con, nil
}
