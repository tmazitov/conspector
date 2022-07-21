package hash

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSha256(str string) string {

	h := sha256.New()
	h.Write([]byte(str))
	sha256 := hex.EncodeToString(h.Sum(nil))

	return sha256
}
