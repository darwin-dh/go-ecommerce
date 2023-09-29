package model

import (
	"encoding/json"

	"github.com/google/uuid"
)

// model for user table
type User struct {
	ID        uuid.UUID       `json:"id"`
	Email     string          `json:"email"`
	Password  string          `json:"password"`
	IsAdmin   bool            `json:"is_admin"`
	Deatils   json.RawMessage `json:"details"`
	CreatedAt int64           `json:"created_at"`
	UpdatedAt int64           `json:"updated_at"`
}

// slice of users
type Users []User
