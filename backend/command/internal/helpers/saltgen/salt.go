package saltgen

import (
	"crypto/rand"
	"encoding/hex"
)

type SaltGenerator interface {
	Generate() (string, error)
}

type Salt struct {}

func (m *Salt) Generate() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}
