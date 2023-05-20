package models

import (
	"github.com/google/uuid"
)

type User struct {
    ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
    Email      string    `gorm:"type:varchar(100);unique_index"`
    Password   string    `gorm:"char:varchar(100)"`
    Registered bool      `gorm:"not null; default:false"`
    ResetToken uuid.UUID `gorm:"type:uuid; default: NULL"`
}
