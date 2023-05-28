package handlers

import (
	"gopher-dispatch/api/models"
	"gopher-dispatch/api/services/policies"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetPolicies handler
func GetPolicies(c *gin.Context) {
	policies, err := policies.GetPolicies()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, policies)
}

// GetPolicy handler
func GetPolicy(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid policy ID"})
		return
	}

	policy, err := policies.GetPolicy(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Policy not found"})
		return
	}
	c.JSON(http.StatusOK, policy)
}

// CreatePolicy handler
func CreatePolicy(c *gin.Context) {
	var newPolicy models.Policy
	if err := c.BindJSON(&newPolicy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	policy, err := policies.CreatePolicy(newPolicy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, policy)
}

// UpdatePolicy handler
func UpdatePolicy(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid policy ID"})
		return
	}

	var updatedPolicy models.Policy
	if err := c.BindJSON(&updatedPolicy); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	policy, err := policies.UpdatePolicy(id, updatedPolicy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, policy)
}

// DeletePolicy handler
func DeletePolicy(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid policy ID"})
		return
	}

	if err := policies.DeletePolicy(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Policy deleted successfully"})
}
