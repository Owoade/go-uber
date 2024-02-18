package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func CompareHashedPassword(pass1 string, pass2 string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(pass1), []byte(pass2))

	if err == nil {
		return true
	}

	fmt.Println(err.Error())

	return false

}
