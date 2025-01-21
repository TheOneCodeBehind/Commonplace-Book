package main

import (
	"time"

	"github.com/google/uuid"
	internal "github.com/sakthi-lucia0567/Commonplace-Book/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseUserToUser(dbUser internal.User) User {
	return User{
		ID:        dbUser.ID.Bytes,
		Username:  dbUser.Username,
		Password:  dbUser.Password,
		Email:     dbUser.Email,
		UpdatedAt: dbUser.UpdatedAt.Time,
		CreatedAt: dbUser.CreatedAt.Time,
	}
}
