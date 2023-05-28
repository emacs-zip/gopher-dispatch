package analytics

import (
	"gopher-dispatch/api/models"
	"gopher-dispatch/pkg/db"
	"time"

	"github.com/google/uuid"
)

func RecordPageView(userID uuid.UUID, page string, duration int) error {
    analyticsEntry := models.PageView{
        ID: uuid.New(),
        UserID: userID,
        Page: page,
        TimeStamp: time.Now(),
        Duration: duration,
    }

    if result := db.GetDB().Create(&analyticsEntry); result.Error != nil {
        return result.Error
    }

    return nil
}

func GetUserPageView(userID string) ([]*models.PageView, error) {
    var pageViewEntries []*models.PageView
    if err := db.GetDB().Where("user_id = ?", userID).Find(&pageViewEntries).Error; err != nil {
        return nil, err
    }

    return pageViewEntries, nil
}
