package db

import (
	"fmt"
	"gopher-dispatch/api/models"

	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func init() {
    username := "postgres"
    password := "root@123"
    dbName := "gopher-dispatch"
    dbHost := "localhost"

    dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
    fmt.Println(dbUri)

    conn, err := gorm.Open("postgres", dbUri);
    if err != nil {
        fmt.Print(err)
    }

    db = conn
    db.Debug().AutoMigrate(&models.AnalyticsEntry{}, &models.User{})
}

func GetDB() *gorm.DB {
    return db
}
