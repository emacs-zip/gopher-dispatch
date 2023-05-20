package dto

import "github.com/google/uuid"

type ForgotPasswordMessage struct {
    Address    string    `json:"address"`
    ResetToken uuid.UUID `json:"recovery_code"`
}
