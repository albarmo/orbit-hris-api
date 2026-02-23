package entities

import (
	"github.com/google/uuid"
)

type PayrollItem struct {
    ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
    PayrollID uuid.UUID `gorm:"type:uuid;not null;index" json:"payroll_id"`
    Type      string    `gorm:"type:varchar" json:"type"`
    Name      string    `gorm:"type:varchar" json:"name"`
    Amount    float64   `gorm:"type:numeric(15,2)" json:"amount"`

    Payroll Payroll `gorm:"foreignKey:PayrollID;references:ID" json:"-"`
}

func (PayrollItem) TableName() string {
    return "payroll_items"
}
