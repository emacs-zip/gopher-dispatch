package dto

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtClaims struct {
    UserID uuid.UUID `json:"userID"`
    Email string `json:"email"`
    TenantID uuid.UUID `json:"tenantID"`
    jwt.RegisteredClaims
}
