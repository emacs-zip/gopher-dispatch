package middleware

import (
	"encoding/json"
	"errors"
	"gopher-dispatch/api/models"
	"gopher-dispatch/pkg/db"
	"net/http"
	"strings"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CasbinAttributes struct {
    UserID uuid.UUID
    Attributes []string
}

func loadUserAttributesFromDatabase(userID string) (string, error) {
    var user models.User
    if err := db.GetDB().Preload("UserAttributes.Attribute").First(&user, "id = ?", userID).Error; err != nil {
        return "", err
    }

    // Assuming each user has a single attribute
    // You might need to update this part according to your real logic
    if len(user.UserAttributes) > 0 {
        return user.UserAttributes[0].Attribute.Value, nil
    }

    return "", errors.New("User has no attributes")
}

func loadTenantPoliciesFromDatabase(enforcer *casbin.Enforcer, tenantID string) error {
    policies := []models.Policy{}
    err := db.GetDB().Where("tenant_id = ?", tenantID).Find(&policies).Error
    if err != nil {
        return err
    }

    for _, policy := range policies {
        enforcer.AddPolicy(policy.Subject, policy.Object, policy.Effect)
    }

    return nil
}

func AttributesMatch(arguments ...interface{}) (interface{}, error) {
    if len(arguments) != 2 {
        return false, errors.New("the number of arguments should be 2")
    }

    requestAttr, ok := arguments[0].(map[string]string)
    if !ok {
        return false, errors.New("failed to assert requestAttr as map[string]string")
    }

    policyAttrJSON, ok := arguments[1].(string)
    if !ok {
        return false, errors.New("failed to assert policyAttrJSON as string")
    }

    var policyAttr map[string]string
    err := json.Unmarshal([]byte(policyAttrJSON), &policyAttr)
    if err != nil {
        return false, err
    }

    for k, v := range policyAttr {
        if requestAttr[k] != v {
            return false, nil
        }
    }

    return true, nil
}

func ABAC() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Extract user ID and tenant ID from the request
        userID := c.GetString("userID")
        tenantID := c.GetString("tenantID")

        // Load user's attributes from the database
        attrs, err := loadUserAttributesFromDatabase(userID)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        // Load policies from the database
        enforcer := db.GetEnforcer()

        // Add the AttributesMatch function to the enforcer
        enforcer.AddFunction("AttributesMatch", AttributesMatch)

        err = loadTenantPoliciesFromDatabase(enforcer, tenantID)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        obj := c.Request.URL.Path
        basePath := "/" + strings.Split(strings.TrimPrefix(obj, "/"), "/")[0]
        sub := attrs

        ok, err := enforcer.Enforce(sub, basePath)
        if err != nil {
            c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        if ok {
            c.Next()
        } else {
            c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "Permission denied"})
        }
    }
}
