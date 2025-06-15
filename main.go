package main

import (
	"log"
	"movefiles/config"
	"net/http"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Menjalankan fungsi Copy files
	//helper.CopyFileTask()

	// Menjalankan Cronjob Golang
	c := gocron.NewScheduler(time.Local)
	c.Every(10).Second().Do(config.ReadEnv)
	c.StartBlocking()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "System status success")
	})

	if err := e.Start(":8080"); err != nil {
		log.Println(err.Error())
	}
}
