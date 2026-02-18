package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type AuditLog struct {
	ID         uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	UserID     uuid.UUID      `gorm:"type:uuid" json:"user_id"`
	Action     string         `gorm:"type:varchar;not null" json:"action"`
	Entity     string         `gorm:"type:varchar;not null" json:"entity"`
	EntityID   uuid.UUID      `gorm:"type:uuid" json:"entity_id"`
	OldValues  datatypes.JSON `gorm:"type:jsonb" json:"old_values"`
	NewValues  datatypes.JSON `gorm:"type:jsonb" json:"new_values"`
	IPAddress  string         `gorm:"type:varchar" json:"ip_address"`
	UserAgent  string         `gorm:"type:text" json:"user_agent"`
	RequestID  uuid.UUID      `gorm:"type:uuid" json:"request_id"`
	Source     string         `gorm:"type:varchar" json:"source"`
	Severity   string         `gorm:"type:varchar" json:"severity"`
	CreatedAt  time.Time      `gorm:"type:timestamptz;not null;default:now()" json:"created_at"`
}

func (AuditLog) TableName() string {
	return "audit_logs"
}
