package middleware

import (
	"net/http"
	"gopher-dispatch/api/models"
	"gopher-dispatch/pkg/db"
	"github.com/gin-gonic/gin"
)

func RBAC(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, exists := c.Get("userID")
		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
			return
		}

		user := &models.User{}
		db.GetDB().Preload("Roles").Where("id = ?", userID).First(user)

		for _, role := range user.Roles {
			if role.Name == requiredRole {
				c.Next()
				return
			}
		}

		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "insufficient permissions"})
	}
}
