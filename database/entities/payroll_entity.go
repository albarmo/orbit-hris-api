package entities

import (
	"time"

	"github.com/google/uuid"
)

type PayrollPeriod struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Month     int       `gorm:"type:int" json:"month"`
	Year      int       `gorm:"type:int" json:"year"`
	StartDate time.Time `gorm:"type:date" json:"start_date"`
	EndDate   time.Time `gorm:"type:date" json:"end_date"`
	IsClosed  bool      `gorm:"default:false" json:"is_closed"`
}

func (PayrollPeriod) TableName() string {
	return "payroll_periods"
}

type Payroll struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	EmployeeID      uuid.UUID `gorm:"type:uuid" json:"employee_id"`
	PayrollPeriodID uuid.UUID `gorm:"type:uuid" json:"payroll_period_id"`
	BasicSalary     float64   `gorm:"type:numeric(15,2)" json:"basic_salary"`
	TotalAllowance  float64   `gorm:"type:numeric(15,2)" json:"total_allowance"`
	TotalDeduction  float64   `gorm:"type:numeric(15,2)" json:"total_deduction"`
	NetSalary       float64   `gorm:"type:numeric(15,2)" json:"net_salary"`
	GeneratedAt     time.Time `gorm:"type:timestamptz;default:now()" json:"generated_at"`

	Employee      Employee      `gorm:"foreignKey:EmployeeID;references:ID" json:"employee"`
	PayrollPeriod PayrollPeriod `gorm:"foreignKey:PayrollPeriodID;references:ID" json:"payroll_period"`
}
