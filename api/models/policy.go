package models

import "github.com/google/uuid"

type Policy struct {
    ID       uuid.UUID `gorm:"type:uuid;primary_key;"`
    Subject  string    `gorm:"type:varchar(100)"`
    Object   string    `gorm:"type:varchar(100)"`
    Effect   string    `gorm:"type:varchar(100)"`
    TenantID uuid.UUID `gorm:"type:uuid"`
    Tenant   Tenant    `gorm:"foreignkey:TenantID"`
}
