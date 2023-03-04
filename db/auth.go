package db

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	Mail     string
	Password string
}

func (a *Auth) HashPassword() (string, error) {
	password := a.Password
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (a *Auth) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func (a *Auth) GenerateRandomHash() string {
	return uuid.New().String()
}
