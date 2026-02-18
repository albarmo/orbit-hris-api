package entities

import "github.com/google/uuid"

type Permission struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"type:varchar;unique;not null" json:"name"`
	Description string    `gorm:"type:varchar" json:"description"`
}
