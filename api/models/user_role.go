package models

import (
	"github.com/google/uuid"
)

type UserRole struct {
    UserId uuid.UUID `gorm:"type:uuid;primary_key"`
    RoleId uuid.UUID `gorm:"type:uuid;primary_key"`
}
