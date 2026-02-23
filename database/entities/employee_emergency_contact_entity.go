package entities

import (
	"time"

	"github.com/google/uuid"
)

type EmployeeEmergencyContact struct {
    ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
    EmployeeID uuid.UUID `gorm:"type:uuid;not null;index" json:"employee_id"`
    Name       string    `gorm:"type:varchar" json:"name"`
    Relationship string  `gorm:"type:varchar" json:"relationship"`
    Phone      string    `gorm:"type:varchar" json:"phone"`
    CreatedAt  time.Time `gorm:"type:timestamp with time zone;default:now()" json:"created_at"`

    Employee Employee `gorm:"foreignKey:EmployeeID;references:ID" json:"-"`
}

func (EmployeeEmergencyContact) TableName() string {
    return "employee_emergency_contacts"
}
