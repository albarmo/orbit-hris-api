package entities

import (
	"time"

	"github.com/google/uuid"
)

type Shift struct {
	ID                   uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name                 string    `gorm:"type:varchar" json:"name"`
	StartTime            time.Time `gorm:"type:time" json:"start_time"`
	EndTime              time.Time `gorm:"type:time" json:"end_time"`
	LateToleranceMinutes int       `gorm:"type:int" json:"late_tolerance_minutes"`
}

func (Shift) TableName() string {
	return "shifts"
}

type EmployeeShift struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	EmployeeID    uuid.UUID `gorm:"type:uuid;not null;index" json:"employee_id"`
	ShiftID       uuid.UUID `gorm:"type:uuid;not null;index" json:"shift_id"`
	EffectiveDate time.Time `gorm:"type:date" json:"effective_date"`

	Employee Employee `gorm:"foreignKey:EmployeeID;references:ID" json:"-"`
	Shift    Shift    `gorm:"foreignKey:ShiftID;references:ID" json:"-"`
}

func (EmployeeShift) TableName() string {
	return "employee_shifts"
}
