package entities

import (
	"time"

	"github.com/google/uuid"
)

type PasswordReset struct {
    ID         uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
    UserID     uuid.UUID  `gorm:"type:uuid;not null;index" json:"user_id"`
    ResetToken uuid.UUID  `gorm:"type:uuid;uniqueIndex" json:"reset_token"`
    ExpiresAt  time.Time  `gorm:"type:timestamptz" json:"expires_at"`
    UsedAt     *time.Time `gorm:"type:timestamptz" json:"used_at"`
    CreatedAt  time.Time  `gorm:"type:timestamptz;default:now()" json:"created_at"`

    User User `gorm:"foreignKey:UserID;references:ID" json:"-"`
}

func (PasswordReset) TableName() string {
    return "password_resets"
}
