package dto

type PermissionBaseDTO struct {
	Name string `json:"name" validate:"required"`
}

// Write Structs

type PermissionCreateDTO struct {
	PermissionBaseDTO
}

type PermissionUpdateDTO struct {
	Name string `json:"name,omitempty" validate:"omitempty"`
}

// Read Structs

type PermissionReadDTO struct {
	PermissionBaseDTO
	ID int64 `json:"id"`
}
