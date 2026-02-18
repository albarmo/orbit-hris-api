package entities

import (
	"time"

	"github.com/google/uuid"
)

type Attendance struct {
	ID           uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	EmployeeID   uuid.UUID  `gorm:"type:uuid" json:"employee_id"`
	LocationID   uuid.UUID  `gorm:"type:uuid" json:"location_id"`
	CheckInTime  time.Time  `gorm:"type:timestamptz" json:"check_in_time"`
	CheckOutTime *time.Time `gorm:"type:timestamptz" json:"check_out_time"`
	Status       string     `gorm:"type:varchar" json:"status"`
	CreatedAt    time.Time  `gorm:"type:timestamp with time zone;default:now()" json:"created_at"`

	Employee Employee `gorm:"foreignKey:EmployeeID;references:ID" json:"employee"`
	Location Location `gorm:"foreignKey:LocationID;references:ID" json:"location"`
}

func (Attendance) TableName() string {
	return "attendance"
}
