package dto

type ForgotPasswordModel struct {
    Email string `json:"email" binding:"required"`
}
