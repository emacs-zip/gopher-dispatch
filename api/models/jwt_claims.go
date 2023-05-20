package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtClaims struct {
    ID uuid.UUID `json:"user_id"`
    Email string `json:"email"`
    jwt.RegisteredClaims
}
