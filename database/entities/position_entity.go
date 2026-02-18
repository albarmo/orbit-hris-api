package entities

import "github.com/google/uuid"

type Position struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	DepartmentID uuid.UUID `gorm:"type:uuid" json:"department_id"`
	Name         string    `gorm:"type:varchar;not null" json:"name"`
	Level        string    `gorm:"type:varchar" json:"level"`

	Department Department `gorm:"foreignKey:DepartmentID;references:ID" json:"department"`
}
