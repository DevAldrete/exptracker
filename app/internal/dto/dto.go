// Package dto: helpers functions for converting models into safe responses to end users
package dto

import (
	"time"

	"github.com/devaldrete/exptrack/app/internal/db"
)

type UserResponse struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func UserToResponse(user db.User) UserResponse {
	res := UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Time,
		UpdatedAt: &user.UpdatedAt.Time,
	}

	return res
}
