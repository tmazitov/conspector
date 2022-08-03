package jwt

import "github.com/gin-gonic/gin"

func (s *Storage) RefreshTokenPair(c *gin.Context, oldTokenPair map[string]string) (map[string]string, error) {

	payload, err := s.verifyToken(c, "active:refresh", oldTokenPair["refresh"])
	if err != nil {
		return nil, ErrInvalidToken
	}

	newTokenPair, err := s.CreateTokenPair(c, payload.Username, payload.UID)
	if err != nil {
		return nil, err
	}

	s.DeleteTokenPair(c, oldTokenPair)

	return newTokenPair, nil
}
