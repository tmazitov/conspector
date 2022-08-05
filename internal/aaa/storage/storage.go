package storage

import (
	"database/sql"
	"log"

	"github.com/tmazitov/conspektor_backend.git/internal/aaa/storage/user"
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

	log.Println("aaa : storage success")

	return &storage
}
