package auth

import (
	"golang.org/x/crypto/bcrypt"
)

func ComparePassword(hashedPassword string, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
