package handlers

import (
	"gopher-dispatch/api/models/dto/request"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateRole(c *gin.Context) {
    var request dto.CreateRoleModel
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    //if err := analyticsService.RecordPageView(request.UserID, request.Page, request.Duration); err != nil {
        //c.JSON(http.StatusInternalServerError, gin.H{"error": "Error recording analytics data"})
    //}

    c.JSON(http.StatusOK, gin.H{})
}

