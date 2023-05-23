package models

import (
	"time"

	"github.com/google/uuid"
)

type PageViewEntry struct {
    Id        uuid.UUID `gorm:"type:uuid;primary_key;"`
    UserId    uuid.UUID `gorm:"type:uuid;"`
    Page      string    `gorm:"type:varchar(100)"`
    TimeStamp time.Time `gorm:"type:time"`
    Duration  int       `gorm:"not null"`
}
