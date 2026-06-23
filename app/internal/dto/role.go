package dto

type RoleBaseDTO struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
}

// Write Structs

type RoleCreateDTO struct {
	RoleBaseDTO
}

type RoleUpdateDTO struct {
	Name        *string `json:"name,omitempty" validate:"omitempty"`
	Description *string `json:"description,omitempty"`
}

// Read Structs

type RoleReadDTO struct {
	RoleBaseDTO
	ID int64 `json:"id"`
}
