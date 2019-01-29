package main

import (
	"golang.org/x/crypto/bcrypt"
)

func hashPassword(plainText string) (encrypted string, err error) {
	bytePass, err := bcrypt.GenerateFromPassword([]byte(plainText), 14)
	if err != nil {
		return
	}

	encrypted = string(bytePass)
	return
}

func comparePassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
