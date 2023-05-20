package main

import (
	"gopher-dispatch/api/routes"
	"gopher-dispatch/pkg/db"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
    file, err := os.OpenFile(time.Now().Format("yyyy-MM-dd"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
    if err != nil {
        log.Fatal(err)
    }

    log.New(file, "INFO: Instantiating database connection", log.Ldate|log.Ltime|log.Lshortfile)
    db := db.GetDB()
    defer db.Close()

    log.New(file, "INFO: Propogating api routes", log.Ldate|log.Ltime|log.Lshortfile)
    r := gin.Default()
    routes.SetupRouter(r)

    log.New(file, "INFO: Starting gopher-dispatch", log.Ldate|log.Ltime|log.Lshortfile)
    r.Run()
}
