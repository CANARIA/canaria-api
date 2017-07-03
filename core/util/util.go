package util

import (
	"fmt"

	"github.com/CANARIA/canaria-api/core/config"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"
)

func ToCrypt(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(hashedPassword)
}

func IsValidPassword(accountPassword string, claimedPassword string) (bool, error) {
	// decodedByte, err := hex.DecodeString(accountPassword)
	// if err != nil {
	// 	fmt.Errorf("failed: decode hashed password", err.Error())
	// 	return false, err
	// }
	// fmt.Println("claimedPassword:", claimedPassword)
	// fmt.Println("accountPassword:", accountPassword)

	if err := bcrypt.CompareHashAndPassword([]byte(accountPassword), []byte(claimedPassword)); err != nil {
		return false, err
	}

	return true, nil

}

func Test(password string) (string, error) {
	dk, err := scrypt.Key([]byte(password), []byte(config.SALT), 16384, 8, 1, 32)
	if err != nil {
		return "", err
	}
	fmt.Println("dk => ", dk)
	return password, nil
}

func GenerateToken() string {
	uuid := uuid.NewV4()
	fmt.Printf("UUIDv4: %s\n", uuid)
	return uuid.String()
}
