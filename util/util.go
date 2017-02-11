package util

import (
	"encoding/hex"

	"fmt"

	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func ToHash(password string) string {
	converted, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return hex.EncodeToString(converted[:])
}

func GenerateToken() string {
	uuid := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", uuid)
	return uuid.String()
}
