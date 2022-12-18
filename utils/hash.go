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

func CompareHash(hash, payloads string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(payloads))

	if err != nil {
		return err
	}

	return nil
}
