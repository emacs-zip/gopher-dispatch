package models

import "github.com/google/uuid"

type Tenant struct {
    ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
    Name      string    `gorm:"type:varchar(100);unique_index"`
    UserLimit int       `gorm:"type:int4"`
    Users     []User    `gorm:"foreignkey:TenantID"`
    Policies  []Policy  `gorm:"foreignkey:TenantID"`
}
