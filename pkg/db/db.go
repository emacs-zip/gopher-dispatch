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


    conn, err := gorm.Open("postgres", dbUri);
    if err != nil {
        fmt.Print(err)
    }

    db = conn
    db.Debug().AutoMigrate(&models.PageViewEntry{}, &models.User{}, &models.Role{}, &models.UserRole{})

    // Create a default role entry "user"
    SetupDefaultRoles()
}

func SetupDefaultRoles() {
    var count int
    db.Model(&models.Role{}).Where("name = ?", "user").Count(&count)
    if count == 0 {
        userRole := models.Role{Name: "user"}
        db.Create(&userRole)
    }
}

func GetDB() *gorm.DB {
    return db
}
