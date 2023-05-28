package request

import "github.com/google/uuid"

type RecordPageViewModel struct {
    UserID   uuid.UUID `json:"userID" binding:"required"`
    Page     string    `json:"page" binding:"required"`
    Duration int       `json:"duration" binding:"required"`
}
