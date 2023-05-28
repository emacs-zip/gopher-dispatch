package models

import (
	"time"

	"github.com/google/uuid"
)

type PageView struct {
    ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
    UserID    uuid.UUID `gorm:"type:uuid;"`
    Page      string    `gorm:"type:varchar(100)"`
    TimeStamp time.Time `gorm:"type:time"`
    Duration  int       `gorm:"not null"`
}
