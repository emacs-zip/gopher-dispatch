package models

import (
	"github.com/google/uuid"
)

type User struct {
    Id         uuid.UUID  `gorm:"type:uuid;primary_key;"`
    Email      string     `gorm:"type:varchar(100);unique_index"`
    Password   string     `gorm:"type:varchar(100)"`
    Roles      []Role     `gorm:"many2many:user_roles;association_foreignkey:ID;foreignkey:ID"`
    Registered bool       `gorm:"not null; default:false"`
    ResetToken uuid.UUID  `gorm:"type:uuid; default: NULL"`
}
