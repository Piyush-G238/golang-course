package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(userPassword *string) {

	hashedPassword, _err := bcrypt.GenerateFromPassword([]byte(*userPassword), 14)
	if _err != nil {
		panic("failed to hashed user password")
	}

	*userPassword = string(hashedPassword)
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
