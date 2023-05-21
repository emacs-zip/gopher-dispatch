package authenticationService

import (
	//"context"
	//"encoding/json"
	"errors"
	"fmt"
	"gopher-dispatch/api/models"
	//"gopher-dispatch/api/models/dto/kafka"
	"gopher-dispatch/pkg/db"
	//"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	//"github.com/segmentio/kafka-go"
	"golang.org/x/crypto/bcrypt"
)

// temp secret, will be handled through env var later. docker, proper deployment, yada
var secret = []byte("SuperCoolSecret!11!!")

func generateToken(user *models.User) (string, error) {
    claims := models.JwtClaims{
        ID: user.ID,
        Email: user.Email,
        RegisteredClaims: jwt.RegisteredClaims{
        	Issuer:    "",
        	Subject:   "",
        	Audience:  []string{},
        	ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
        	NotBefore: jwt.NewNumericDate(time.Now()),
        	IssuedAt:  jwt.NewNumericDate(time.Now()),
        	ID:        "",
        },
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    tokenString, err := token.SignedString(secret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func SignInWithEmail(email string, password string) (string, error) {
    user := &models.User{}
    if err := db.GetDB().Where("email = ?", email).First(user).Error; err != nil {
        return "", err
    }

    if !user.Registered {
        return "", errors.New("user must be registered to login")
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
        return "", errors.New("invalid login credentials")
    }

    token, err := generateToken(user)
    if err != nil {
        return "", err
    }

    return token, nil
}

func SignInWithJwt(token string) error  {
    // Yucky interface{}, but that's the docs
	decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})

    if err != nil {
        return fmt.Errorf("invalid JWT token: %w", err)
    }

	if claims, ok := decodedToken.Claims.(jwt.MapClaims); ok && decodedToken.Valid {
		userID := claims["user_id"].(string)
		user := &models.User{}
		if err := db.GetDB().Where("id = ?", userID).First(user).Error; err != nil {
			return err
		}

		return nil
	}

    return nil
}

func SignUp(email string, password string) (*models.User, error) {
    hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

    user := models.User{
        ID: uuid.New(),
        Email: email,
        Password: string(hashedPassword),
    }

    if result := db.GetDB().Create(&user); result.Error != nil {
        return nil, result.Error
    }

    // Begin kafka message to bucktooth-envoy.registration
    // Leaving commented until email service implemented
    /*writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "bucktooth-envoy.forgotpassword",
    })

    registrationMessage := dto.RegistrationMessage{
        Address: user.Email,
    }

    registrationMessageBytes, err := json.Marshal(registrationMessage)
    if err != nil {
        log.Fatal("Failed to encode email message: ", err)
        return nil, err
    }

    err = writer.WriteMessages(context.Background(), kafka.Message{
        Value: registrationMessageBytes,
    })

    if err != nil {
        log.Fatal("Failed to write message: ", err)
        return nil, err
    }

    writer.Close()*/

    return &user, nil
}

func ForgotPassowrd(email string) error  {
    user := &models.User{}
    if err := db.GetDB().Where("email = ?", email).First(user).Error; err != nil {
        return err
    }

    resetToken := uuid.New()
    user.ResetToken = resetToken

    if err := db.GetDB().Save(user).Error; err != nil {
        return err;
    }

    // Begin kafka message to bucktooth-envoy.forgotpassword
    // Leaving commented until email service implemented
    /*writer := kafka.NewWriter(kafka.WriterConfig{
        Brokers: []string{"localhost:9092"},
        Topic:   "bucktooth-envoy.forgotpassword",
    })

    emailMessage := dto.ForgotPasswordMessage{
        Address: email,
        ResetToken: resetToken,
    }

    emailMessageBytes, err := json.Marshal(emailMessage)
    if err != nil {
        log.Fatal("Failed to encode email message: ", err)
        return err
    }

    err = writer.WriteMessages(context.Background(), kafka.Message{
        Value: emailMessageBytes,
    })

    if err != nil {
        log.Fatal("Failed to write message: ", err)
        return err
    }

    writer.Close()*/

    return nil
}
