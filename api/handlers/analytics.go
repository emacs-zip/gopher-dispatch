package handlers

import (
	"gopher-dispatch/api/models/dto/request"
	"gopher-dispatch/api/services/analytics"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RecordPageView(c *gin.Context) {
    var request dto.RecordPageViewModel
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := analyticsService.RecordPageView(request.UserID, request.Page, request.Duration); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error recording analytics data"})
    }

    c.JSON(http.StatusOK, gin.H{})
}

func GetUserPageView(c *gin.Context) {
    userID := c.Param("user_id")

    pageViewData, err := analyticsService.GetUserPageView(userID)

    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to get user page view data"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"page_view_data": pageViewData})
}
