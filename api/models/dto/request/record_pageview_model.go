package dto

import "github.com/google/uuid"

type RecordPageViewModel struct {
    UserId   uuid.UUID `json:"userId" binding:"required"`
    Page     string    `json:"page" binding:"required"`
    Duration int       `json:"duration" binding:"required"`
}
