package authentication

import (
	//"context"
	//"encoding/json"
	"errors"
	"fmt"
	"gopher-dispatch/api/models"
	"gopher-dispatch/api/models/dto"
	"strings"

	//"gopher-dispatch/api/models/dto/kafka"
	"gopher-dispatch/pkg/config"
	"gopher-dispatch/pkg/db"

	//"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	//"github.com/segmentio/kafka-go"
	"golang.org/x/crypto/bcrypt"
)

func createUser(email string, password string, tenantID uuid.UUID, admin bool) error {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }

    // Create new user
    userID := uuid.New()
    user := &models.User {
        ID: userID,
        Email: email,
        Password: string(hashedPassword),
        Registered: true, // TODO: Default to false, send kafka message to envoy
        TenantID: tenantID,
    }

    if err := db.GetDB().Create(&user).Error; err != nil {
        fmt.Println("user error")
        fmt.Println(err)
        return err
    }

    defaultAttr := "user"
    if admin {
        defaultAttr = "admin"
    }

    // Find default role attribute
    roleAttr := &models.Attribute{}
    if err := db.GetDB().Where("value = ?", defaultAttr).First(&roleAttr).Error; err != nil {
        return err
    }

    // Assign default attribute to user
    userAttr := &models.UserAttribute{
        ID: uuid.New(),
        UserID: userID,
        AttributeID: roleAttr.ID,
    }

    if err := db.GetDB().Create(&userAttr).Error; err != nil {
        return err
    }

    return nil
}

func generateToken(user *models.User) (string, error) {
    claims := dto.JwtClaims{
        UserID: user.ID,
        Email: user.Email,
        TenantID: user.TenantID,
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

    tokenString, err := token.SignedString([]byte(config.Get().JWTSecret))
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

func SignInWithJwt(authHeader string) error  {
        const BearerSchema = "Bearer "

        // check if Authorization header is correctly formatted
        if !strings.HasPrefix(authHeader, BearerSchema) {
            return nil
        }

        token := authHeader[len(BearerSchema):]

        decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
            }

            return []byte(config.Get().JWTSecret), nil
        })

    if err != nil {
        return fmt.Errorf("invalid JWT token: %w", err)
    }

	if claims, ok := decodedToken.Claims.(jwt.MapClaims); ok && decodedToken.Valid {
		userID := claims["userID"].(string)
		user := &models.User{}
		if err := db.GetDB().Where("id = ?", userID).First(user).Error; err != nil {
			return err
		}

		return nil
	}

    return nil
}

func SignUp(email string, password string) error {
    // Create a tenant
    tenant := &models.Tenant{
        ID: uuid.New(),
        Name: email,
        UserLimit: 10,
    }

    if err := db.GetDB().Create(&tenant).Error; err != nil {
        return err
    }

    err := createUser(email, password, tenant.ID, true)
    if err != nil {
        return err
    }

    // Create default policy linked to new tenant
    policy := models.Policy{
        ID: uuid.New(),
        Subject: "admin",
        Object: "/page-view",
        Effect: "allow",
        TenantID: tenant.ID,
    }

    if err := db.GetDB().Create(&policy).Error; err != nil {
        return err
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

    return nil
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
