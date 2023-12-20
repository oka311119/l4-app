package helpers

import (
	"crypto/rand"
	"encoding/hex"
)

type ISaltGenerator interface {
	Generate() (string, error)
}

type SaltGenerator struct {}

func NewSaltGenerator() *SaltGenerator {
	return &SaltGenerator{}
}

func (m *SaltGenerator) Generate() (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}
