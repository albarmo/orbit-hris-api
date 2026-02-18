package entities

import "github.com/google/uuid"

// Explicit join table for User-Role relationship
type UserRole struct {
	UserID uuid.UUID `gorm:"primaryKey"`
	RoleID uuid.UUID `gorm:"primaryKey"`
}

func (UserRole) TableName() string {
	return "user_roles"
}

// Explicit join table for Role-Permission relationship
type RolePermission struct {
	RoleID       uuid.UUID `gorm:"primaryKey"`
	PermissionID uuid.UUID `gorm:"primaryKey"`
}

func (RolePermission) TableName() string {
	return "role_permissions"
}
