package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type User struct {
    gorm.Model
    ID         uuid.UUID `gorm:"type:uuid;primary_key;"`
    Registered bool      `gorm:"not null; default:false"`
    Email      string    `gorm:"type:varchar(100);unique_index"`
    Password   string    `gorm:"char:varchar(100)"`
    ResetToken uuid.UUID `gorm:"type:uuid"`
}
