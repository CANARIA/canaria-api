package util

import (
	"encoding/hex"

	"golang.org/x/crypto/bcrypt"
)

func ToHash(password string) string {
	converted, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return hex.EncodeToString(converted[:])
}
