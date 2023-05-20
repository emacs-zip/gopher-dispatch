package main

import (
	"gopher-dispatch/api/routes"
	"gopher-dispatch/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
    // TODO: logrus

    db := db.GetDB()
    defer db.Close()

    router := gin.Default()
    routes.SetupRouter(router)

    router.Run()
}
