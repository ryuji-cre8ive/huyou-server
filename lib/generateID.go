package lib

import (
	"github.com/google/uuid"
)

func GenerateRandomHash() string {
	return uuid.New().String()
}