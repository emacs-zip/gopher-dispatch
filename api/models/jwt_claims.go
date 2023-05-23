package models

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtClaims struct {
    Id uuid.UUID `json:"userId"`
    Email string `json:"email"`
    jwt.RegisteredClaims
}
