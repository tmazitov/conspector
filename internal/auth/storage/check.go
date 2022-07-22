package storage

import (
	"database/sql"
	"errors"
)

var (
	ErrUsernameIsNotUnique = errors.New("invalid credentials")
	ErrEmailIsNotUnique    = errors.New("invalid credentials")
	ErrPasswordIsInvalid   = errors.New("invalid credentials")
)

func (s *Storage) CheckUsername(username string) error {
	var (
		execResult *sql.Row
		isExists   bool
		execString string
		err        error
	)

	execString = "select exists(select 1 from users where username = $1)"
	execResult = s.conn.QueryRow(execString, username)

	if err = execResult.Scan(&isExists); err != nil {
		return err
	}

	if isExists {
		return ErrUsernameIsNotUnique
	}

	return nil
}

func (s *Storage) CheckEmail(email string) error {
	var (
		execResult *sql.Row
		isExists   bool
		execString string
		err        error
	)

	execString = "select exists(select 1 from users where email = $1)"
	execResult = s.conn.QueryRow(execString, email)

	if err = execResult.Scan(&isExists); err != nil {
		return err
	}

	if isExists {
		return ErrEmailIsNotUnique
	}

	return nil
}

func (s *Storage) CheckPassword(password string, username string) error {
	var (
		execResult      *sql.Row
		storagePassword string
		execString      string
		err             error
	)

	execString = "select password from users where username = $1"
	execResult = s.conn.QueryRow(execString, username)

	if err = execResult.Scan(&storagePassword); err != nil {
		return err
	}

	if storagePassword != password {
		return ErrPasswordIsInvalid
	}

	return nil
}
