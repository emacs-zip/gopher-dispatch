package models

import (
	"github.com/google/uuid"
)

type UserRole struct {
    UserID uuid.UUID `gorm:"type:uuid;primary_key"`
    RoleID uuid.UUID `gorm:"type:uuid;primary_key"`
}
