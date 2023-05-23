package models

import (
    "github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Role struct {
    Id       uuid.UUID  `gorm:"type:uuid;primary_key"`
    Name     string     `gorm:"type:varchar(30)"`
    UserRole []UserRole `gorm:"foreignKey:RoleId;references:Id"`
}

func (role *Role) BeforeCreate(scope *gorm.Scope) error {
    uuid := uuid.New()
    return scope.SetColumn("Id", uuid)
}
