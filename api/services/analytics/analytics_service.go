package analytics

import (
	"gopher-dispatch/api/models"
	"gopher-dispatch/pkg/db"
	"time"

	"github.com/google/uuid"
)

func RecordPageView(userId uuid.UUID, page string, duration int) error {
    analyticsEntry := models.PageView{
        Id: uuid.New(),
        UserId: userId,
        Page: page,
        TimeStamp: time.Now(),
        Duration: duration,
    }

    if result := db.GetDB().Create(&analyticsEntry); result.Error != nil {
        return result.Error
    }

    return nil
}

func GetUserPageView(userId string) ([]*models.PageView, error) {
    var pageViewEntries []*models.PageView
    if err := db.GetDB().Where("user_id = ?", userId).Find(&pageViewEntries).Error; err != nil {
        return nil, err
    }

    return pageViewEntries, nil
}
