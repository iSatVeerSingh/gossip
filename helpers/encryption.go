package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, bool) {
	hp, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", false
	}

	return string(hp), true
}

func ValidatePassword(password string, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
