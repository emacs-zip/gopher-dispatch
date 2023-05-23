package dto

type SignUpModel struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type SignInModel struct {
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

type ForgotPasswordModel struct {
    Email string `json:"email" binding:"required"`
}
