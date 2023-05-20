package services

import (
	"gopher-dispatch/api/models"
	"time"

	"github.com/google/uuid"
	"github.com/golang-jwt/jwt/v5"
)

type JwtClaims struct {
    ID uuid.UUID `json:"user_id"`
    Email string `json:"email"`
    jwt.RegisteredClaims
}

func generateToken(user *models.User) (string, error) {
    // temp secret, will be handled through env var later. docker, proper deployment, yada
    secret := []byte("SuperCoolSecret!11!!")

    claims := JwtClaims{
        user.ID,
        user.Email,
        jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(secret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func SignInWithEmailService(email string, password string) (string, string, error) {
    // TODO: Implement functionality

    return "", "", nil
}

func SignInWithJwtService(token string) (*models.User, error)  {
    // TODO: Implement functionality

    return &models.User{}, nil
}

func SignUpService(user *models.User) error  {
    // TODO: Implement functionality

    return nil
}

func ForgotPassowrdService(email string) error  {
    // TODO: Implement functionality

    return nil
}
