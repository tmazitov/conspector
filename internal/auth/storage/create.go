package storage

import (
	"github.com/tmazitov/conspektor_backend.git/internal/auth/dto"
)

func (s *Storage) UserCreate(user dto.CreateUser) error {
	execString := "insert into Users (username, password, email, uid) values ($1, $2, $3, $4)"
	_, err := s.conn.Exec(execString, user.Username, user.Password, user.Email, user.UID)
	return err
}
