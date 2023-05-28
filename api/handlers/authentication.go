package handlers

import (
	"fmt"
	"gopher-dispatch/api/models/dto/request"
	"gopher-dispatch/api/services/authentication"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
    var request request.SignUpModel
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := authentication.SignUp(request.Email, request.Password)

    if err != nil {
        fmt.Printf("ERROR: %s", err)
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error signing up user"})
        return
    }

    c.JSON(http.StatusOK, gin.H{})
}

func SignInWithEmail(c *gin.Context) {
    var request request.SignInModel
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    token, err := authentication.SignInWithEmail(request.Email, request.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func SignInWithJwt(c *gin.Context) {
	token := c.GetHeader("Authorization")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization header is missing"})
		return
	}

	err := authentication.SignInWithJwt(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

func ForgotPassword(c *gin.Context) {
    var request request.ForgotPasswordModel
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err := authentication.ForgotPassowrd(request.Email)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Email does not exist"})
    }

    c.JSON(http.StatusOK, gin.H{"message": "Reset token has been sent to your email"})
}
