package dto

type UsersRoleBaseDTO struct {
	UserID   int64 `json:"user_id" validate:"required"`
	RoleID   int64 `json:"role_id" validate:"required"`
	IsActive bool  `json:"is_active"`
}

// Write Structs

type UsersRoleCreateDTO struct {
	UsersRoleBaseDTO
}

type UsersRoleUpdateDTO struct {
	IsActive *bool `json:"is_active,omitempty"`
}

// Read Structs

type UsersRoleReadDTO struct {
	UsersRoleBaseDTO
}
