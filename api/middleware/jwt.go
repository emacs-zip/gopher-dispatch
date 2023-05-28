package middleware

import (
	"fmt"
	"gopher-dispatch/pkg/config"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        const BearerSchema = "Bearer "
        authHeader := c.GetHeader("Authorization")

        // check if Authorization header is correctly formatted
        if !strings.HasPrefix(authHeader, BearerSchema) {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing Authorization header"})
            return
        }

        token := authHeader[len(BearerSchema):]

        fmt.Println("Authorization header: ", authHeader)

        decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
            }

            return []byte(config.Get().JWTSecret), nil
        })

        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        if claims, ok := decodedToken.Claims.(jwt.MapClaims); ok && decodedToken.Valid {
            c.Set("userID", claims["userID"].(string))
            c.Set("tenantID", claims["tenantID"].(string))
        } else {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            return
        }

        c.Next()
    }
}
