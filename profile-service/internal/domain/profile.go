package domain

import (
	"time"

	"github.com/google/uuid"
)

type Profile struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	Bio       string    `json:"bio"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProfile(email, name string) *Profile {
	return &Profile{
		ID:        uuid.NewString(),
		Email:     email,
		Name:      name,
		Bio:       "No Bio",
		CreatedAt: time.Now(),
	}
}
