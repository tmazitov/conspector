package user

import "database/sql"

type UserStorage struct {
	conn *sql.DB
}

func NewUserStorage(conn *sql.DB) *UserStorage {
	return &UserStorage{conn: conn}
}
