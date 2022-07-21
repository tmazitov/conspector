package storage

import "database/sql"

type Storage struct {
	conn *sql.DB
}

func NewStorage(conn *sql.DB) *Storage {

	storage := Storage{conn: conn}

	return &storage
}
