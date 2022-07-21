package storage

import (
	"database/sql"
	"errors"
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
		return errors.New("username is not unique")
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
		return errors.New("email is not unique")
	}

	return nil
}
