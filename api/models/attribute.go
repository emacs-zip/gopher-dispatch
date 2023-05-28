package models

import "github.com/google/uuid"

type Attribute struct {
    ID                 uuid.UUID           `gorm:"type:uuid;primary_key;"`
    Value              string              `gorm:"type:varchar(100);unique_index"`
    UserAttributes     []UserAttribute     `gorm:"foreignkey:AttributeID"`
}
