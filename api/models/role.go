package models

import (
    "github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Role struct {
    ID       uuid.UUID  `gorm:"type:uuid;primary_key"`
    Name     string     `gorm:"type:varchar(30)"`
    UserRole []UserRole `gorm:"foreignKey:RoleID;references:ID"`
}

func (role *Role) BeforeCreate(scope *gorm.Scope) error {
    uuid := uuid.New()
    return scope.SetColumn("ID", uuid)
}
