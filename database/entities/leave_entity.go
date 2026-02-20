package entities

import (
	"time"

	"github.com/google/uuid"
)

type Leave struct {
    ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
    EmployeeID uuid.UUID `gorm:"type:uuid" json:"employee_id"`
    StartDate  time.Time `gorm:"type:date" json:"start_date"`
    EndDate    time.Time `gorm:"type:date" json:"end_date"`
    Reason     string    `gorm:"type:text" json:"reason"`
    Status     string    `gorm:"type:varchar" json:"status"`
    CreatedAt  time.Time `gorm:"type:timestamp with time zone;default:now()" json:"created_at"`

    Employee Employee `gorm:"foreignKey:EmployeeID;references:ID" json:"employee"`
}

func (Leave) TableName() string {
    return "leaves"
}
