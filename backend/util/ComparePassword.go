package util

import "golang.org/x/crypto/bcrypt"

func CompareHashPassword(password string, hash string) (error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}