package dto

import "github.com/google/uuid"

type RecordPageViewModel struct {
    UserID   uuid.UUID `json:"user_id" binding:"required"`
    Page     string    `json:"page" binding:"required"`
    Duration int       `json:"duration" binding:"required"`
}
