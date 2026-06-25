// Package dto: helpers functions for converting models into safe responses to end users
package dto

import "time"

type UserResponse struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
