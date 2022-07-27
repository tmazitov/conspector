package storage

import (
	"database/sql"

	"github.com/tmazitov/conspektor_backend.git/internal/auth/storage/user"
)

type Storage struct {
	conn *sql.DB
	User user.UserStorage
}

func NewStorage(conn *sql.DB) *Storage {

	user := user.NewUserStorage(conn)
	storage := Storage{
		conn: conn,
		User: *user,
	}

	return &storage
}
