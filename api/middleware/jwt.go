package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secret = []byte("SuperCoolSecret!11!!")

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        const BearerSchema = "Bearer "
        authHeader := c.GetHeader("Authorization")

        // check if Authorization header is correctly formatted
        if !strings.HasPrefix(authHeader, BearerSchema) {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid or missing Authorization header"})
            return
        }

        token := authHeader[len(BearerSchema):]

        decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }

            return []byte(secret), nil
        })

        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }

        if claims, ok := decodedToken.Claims.(jwt.MapClaims); ok && decodedToken.Valid {
            c.Set("userID", claims["user_id"].(string))
        } else {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }

        c.Next()
    }
}

/*func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        const BearerSchema = "Bearer "
        authHeader := c.GetHeader("Authorization")
        token := authHeader[len(BearerSchema):]

        decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
            }

            return []byte(secret), nil
        })

        if err != nil {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }

        if claims, ok := decodedToken.Claims.(jwt.MapClaims); ok && decodedToken.Valid {
            c.Set("userID", claims["user_id"].(string))
        } else {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            return
        }

        c.Next()
    }
}*/
