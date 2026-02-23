package entities

import (
	"time"

	"github.com/google/uuid"
)

type AuthToken struct {
    ID         uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
    UserID     uuid.UUID `gorm:"type:uuid;not null;index" json:"user_id"`
    AccessUUID uuid.UUID `gorm:"type:uuid;uniqueIndex" json:"access_uuid"`
    RefreshUUID uuid.UUID `gorm:"type:uuid;uniqueIndex" json:"refresh_uuid"`
    ExpiresAt  time.Time `gorm:"type:timestamptz" json:"expires_at"`
    CreatedAt  time.Time `gorm:"type:timestamptz;default:now()" json:"created_at"`

    User User `gorm:"foreignKey:UserID;references:ID" json:"-"`
}

func (AuthToken) TableName() string {
    return "auth_tokens"
}
