package entities

import "github.com/google/uuid"

type Role struct {
	ID          uuid.UUID    `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name        string       `gorm:"type:varchar;unique;not null" json:"name"`
	Description string       `gorm:"type:varchar" json:"description"`
	Permissions []Permission `gorm:"many2many:role_permissions;" json:"permissions,omitempty"`
}
