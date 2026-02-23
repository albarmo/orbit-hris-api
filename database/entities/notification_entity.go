package entities

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID      uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
	Type        string     `gorm:"type:varchar" json:"type"`
	Title       string     `gorm:"type:varchar" json:"title"`
	Message     string     `gorm:"type:text" json:"message"`
	ReferenceID *uuid.UUID `gorm:"type:uuid" json:"reference_id"`
	IsRead      bool       `gorm:"default:false" json:"is_read"`
	CreatedAt   time.Time  `gorm:"type:timestamptz;not null;default:now()" json:"created_at"`

	User User `gorm:"foreignKey:UserID;references:ID" json:"-"`
}

func (Notification) TableName() string {
	return "notifications"
}
