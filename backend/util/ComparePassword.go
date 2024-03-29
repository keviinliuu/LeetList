package util

import "golang.org/x/crypto/bcrypt"

func CompareHashPassword(hash string, password string) (error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}