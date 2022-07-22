package storage

import (
	"database/sql"

	"github.com/tmazitov/conspektor_backend.git/internal/auth/models"
)

func (s *Storage) GetUserInfo(username string) (models.User, error) {
	var (
		execResult *sql.Row
		user       models.User
		execString string
		err        error
	)

	execString = "select uid, username from users where username = $1"
	execResult = s.conn.QueryRow(execString, username)

	if err = execResult.Scan(&user.UID, &user.Username); err != nil {
		return user, err
	}

	return user, nil
}
