package db

import (
	"fmt"
	"gopher-dispatch/api/models"
	"gopher-dispatch/pkg/config"
	"os"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
    db *gorm.DB
    enforcer *casbin.Enforcer
)

func createDefaultRoleAttributes() error {
    // Create "user" attribute if one doesn't exist
    userRoleAttribute := &models.Attribute{}
    if err := db.Where("value = ?", "user").First(userRoleAttribute).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            userRoleAttribute = &models.Attribute{
                ID: uuid.New(),
                Value: "user",
            }

            if createErr := db.Create(userRoleAttribute).Error; createErr != nil {
                return createErr
            }
        } else {
            // Other database errors
            return err
        }
    }

    // Create "admin" attribute if one doesn't exist
    adminRoleAttribute := &models.Attribute{}
    if err := db.Where("value = ?", "admin").First(adminRoleAttribute).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            adminRoleAttribute = &models.Attribute{
                ID: uuid.New(),
                Value: "admin",
            }

            if createErr := db.Create(adminRoleAttribute).Error; createErr != nil {
                return createErr
            }
        } else {
            // Other database errors
            return err
        }
    }

    return nil
}

func init() {
    // Connect to db
    dbUri := fmt.Sprintf(
        "host=%s user=%s dbname=%s sslmode=disable password=%s",
        config.Get().DBHost,
        config.Get().DBUsername,
        config.Get().DBName,
        config.Get().DBPassword,
    )

    conn, err := gorm.Open("postgres", dbUri);
    if err != nil {
        fmt.Print(err)
    }

    // Create tables
    db = conn
    db.Debug().AutoMigrate(
        &models.Attribute{},
        &models.PageView{},
        &models.Policy{},
        &models.Tenant{},
        &models.User{},
        &models.UserAttribute{},
    )

    if err := createDefaultRoleAttributes(); err != nil {
        fmt.Print(err)
    }

    // Init casper gorm adapter
    a, err := gormadapter.NewAdapter("postgres", dbUri, true)
    if err != nil {
        fmt.Print(err)
        os.Exit(1)
    }

    // Init casper enforcer
    e, err := casbin.NewEnforcer("pkg/config/abac_model.conf", a)
    if err != nil {
        fmt.Print(err)
        os.Exit(1)
    }

    e.LoadPolicy()

    enforcer = e
}

func GetDB() *gorm.DB {
    return db
}

func GetEnforcer() *casbin.Enforcer {
    return enforcer
}
