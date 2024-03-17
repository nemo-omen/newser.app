package usecase

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(p string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(p), 8)
	return string(hashed)
}
