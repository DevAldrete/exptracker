package dto

import (
	"time"
)

type UserBaseDTO struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Write Structs

type UserCreateDTO struct {
	UserBaseDTO
	Password string `json:"password" validate:"required,min=8"`
}

type UserUpdateDTO struct {
	UserBaseDTO
	Password string `json:"password,omitempty" validate:"omitempty,min=8"`
}

// Read Structs

type UserReadDTO struct {
	UserBaseDTO
	ID        int64      `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
