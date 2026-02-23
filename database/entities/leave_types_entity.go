package entities

import (
	"github.com/google/uuid"
)

type LeaveType struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name           string    `gorm:"type:varchar;not null" json:"name"`
	MaxDaysPerYear int       `gorm:"type:int" json:"max_days_per_year"`
	IsPaid         bool      `gorm:"default:false" json:"is_paid"`
}

func (LeaveType) TableName() string {
	return "leave_types"
}

type LeaveBalance struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	EmployeeID    uuid.UUID `gorm:"type:uuid;not null;index" json:"employee_id"`
	LeaveTypeID   uuid.UUID `gorm:"type:uuid;not null;index" json:"leave_type_id"`
	Year          int       `gorm:"type:int" json:"year"`
	RemainingDays int       `gorm:"type:int" json:"remaining_days"`
}

func (LeaveBalance) TableName() string {
	return "leave_balances"
}
