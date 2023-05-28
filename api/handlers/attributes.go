package handlers

import (
	"gopher-dispatch/api/models"
	"gopher-dispatch/api/services/attributes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetAttributes handler
func GetAttributes(c *gin.Context) {
	attributes, err := attributes.GetAttributes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, attributes)
}

// GetAttribute handler
func GetAttribute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attribute ID"})
		return
	}

	attribute, err := attributes.GetAttribute(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attribute not found"})
		return
	}
	c.JSON(http.StatusOK, attribute)
}

// CreateAttribute handler
func CreateAttribute(c *gin.Context) {
	var newAttribute models.Attribute
	if err := c.BindJSON(&newAttribute); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attribute, err := attributes.CreateAttribute(newAttribute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, attribute)
}

// UpdateAttribute handler
func UpdateAttribute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attribute ID"})
		return
	}

	var updatedAttribute models.Attribute
	if err := c.BindJSON(&updatedAttribute); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	attribute, err := attributes.UpdateAttribute(id, updatedAttribute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, attribute)
}

// DeleteAttribute handler
func DeleteAttribute(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attribute ID"})
		return
	}

	if err := attributes.DeleteAttribute(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Attribute deleted successfully"})
}
