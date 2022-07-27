package user

import userDto "github.com/tmazitov/conspektor_backend.git/internal/user/dto/user"

func (s *UserStorage) Create(user userDto.CreateUser) error {
	execString := "insert into Users (username, password, email, uid) values ($1, $2, $3, $4)"
	_, err := s.conn.Exec(execString, user.Username, user.Password, user.Email, user.UID)
	return err
}
