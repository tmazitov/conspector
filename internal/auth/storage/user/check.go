package user

import (
	"database/sql"
	"errors"
	"log"

	"github.com/tmazitov/conspektor_backend.git/internal/auth/dto"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrServer             = errors.New("server error")
)

func (s *UserStorage) CheckUsername(username string) error {
	var (
		execResult *sql.Row
		isExists   bool
		execString string
		err        error
	)

	execString = "select exists(select 1 from users where username = $1)"
	execResult = s.conn.QueryRow(execString, username)

	if err = execResult.Scan(&isExists); err != nil {
		log.Printf("check username : %s", err)
		return ErrServer
	}

	if isExists {
		return ErrInvalidCredentials
	}

	return nil
}

func (s *UserStorage) CheckEmail(email string) error {
	var (
		execResult *sql.Row
		isExists   bool
		execString string
		err        error
	)

	execString = "select exists(select 1 from users where email = $1)"
	execResult = s.conn.QueryRow(execString, email)

	if err = execResult.Scan(&isExists); err != nil {
		log.Printf("check email : %s", err)
		return ErrServer
	}

	if isExists {
		return ErrInvalidCredentials
	}

	return nil
}

func (s *UserStorage) CheckPassword(check dto.CheckPassword) error {
	var (
		execResult      *sql.Row
		storagePassword string
		execString      string
		err             error
	)

	execString = "select password from users where username = $1"
	execResult = s.conn.QueryRow(execString, check.Username)

	if err = execResult.Scan(&storagePassword); err != nil {
		log.Printf("check password : %s", err)
		return ErrServer
	}

	if storagePassword != check.Password {
		return ErrInvalidCredentials
	}

	return nil
}
