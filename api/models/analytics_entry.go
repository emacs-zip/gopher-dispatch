package models

import (
	"github.com/google/uuid"
)

type AnalyticsEntry struct {
    ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
    UserID   uuid.UUID `gorm:"type:uuid;"`
    Page     string    `gorm:"type:varchar(100)"`
    Duration int       `gorm:"not null"`
}
