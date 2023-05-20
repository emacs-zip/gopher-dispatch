package handlers

import (
	"gopher-dispatch/api/models"
	"gopher-dispatch/api/models/dto"
	"gopher-dispatch/api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignInWithEmail(c *gin.Context) {
    var data dto.SignInModel
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user, token, err := services.SignInWithEmailService(data.Email, data.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
    }

    c.JSON(http.StatusOK, gin.H{"data": user, "token": token})
}

func SignInWithJwt(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is missing"})
		return
	}

	user, err := services.SignInWithJwtService(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func SignUp(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := services.SignUpService(&user); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error signing up user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"data": user})
}

func ForgotPassword(c *gin.Context) {
    var data dto.ForgotPasswordModel
    if err := c.ShouldBindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := services.ForgotPassowrdService(data.Email)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email does not exist"})
    }

    c.JSON(http.StatusOK, gin.H{"message": "Reset token has been sent to your email"})
}
