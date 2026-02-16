package domain

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           string
	Email        string
	Role         string
	Name         string
	PasswordHash string
	CreatedAt    time.Time
}

func NewUser(email, role, name, hashPass string) *User {
	return &User{
		ID:           uuid.NewString(),
		Email:        email,
		Role:         role,
		Name:         name,
		PasswordHash: hashPass,
		CreatedAt:    time.Now(),
	}
}
