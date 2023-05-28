package models

import "github.com/google/uuid"

type UserAttribute struct {
    ID          uuid.UUID `gorm:"type:uuid;primary_key;"`
    UserID      uuid.UUID `gorm:"type:uuid"`
    AttributeID uuid.UUID `gorm:"type:uuid"`
    User        User      `gorm:"foreignkey:UserID"`
    Attribute   Attribute `gorm:"foreignkey:AttributeID"`
}
