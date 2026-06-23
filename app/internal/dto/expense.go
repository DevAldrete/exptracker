package dto

import (
	"time"
)

type ExpenseBaseDTO struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description,omitempty"`
	Cents       int64  `json:"cents" validate:"required"`
}

// Write Structs

type ExpenseCreateDTO struct {
	ExpenseBaseDTO
	UserID int8 `json:"user_id"`
}

type ExpenseUpdateDTO struct {
	Name        string `json:"name,omitempty" validate:"omitempty"`
	Description string `json:"description,omitempty"`
	Cents       int64  `json:"cents,omitempty" validate:"omitempty"`
}

// Read Structs

type ExpenseReadDTO struct {
	ExpenseBaseDTO
	ID        int64      `json:"id"`
	UserID    int8       `json:"user_id" validate:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdateAt  *time.Time `json:"updated_at,omitempty"`
}
