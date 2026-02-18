package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Location struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name         string         `gorm:"type:varchar;not null" json:"name"`
	Latitude     float64        `gorm:"type:decimal" json:"latitude"`
	Longitude    float64        `gorm:"type:decimal" json:"longitude"`
	RadiusMeters int            `gorm:"type:int" json:"radius_meters"`
	Polygon      datatypes.JSON `gorm:"type:jsonb" json:"polygon"`
	IsActive     bool           `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time      `gorm:"type:timestamp with time zone;default:now()" json:"created_at"`
}
