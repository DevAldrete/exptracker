package dto

type RolesPermissionBaseDTO struct {
	RoleID       int64 `json:"role_id" validate:"required"`
	PermissionID int64 `json:"permission_id" validate:"required"`
}

// Write Structs

type RolesPermissionCreateDTO struct {
	RolesPermissionBaseDTO
}

// Read Structs

type RolesPermissionReadDTO struct {
	RolesPermissionBaseDTO
}
