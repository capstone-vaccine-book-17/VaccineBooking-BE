package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashBcrypt(password string) (string, error) {
	pw, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(pw), nil
}
