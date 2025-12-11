package main

import (
	"echo-products-api/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.DatabaseInit()
	gorm := config.DB()

	dbGorm, err := gorm.DB()

	if err != nil {
		e.Logger.Fatal("Failed to get generic database object:", err)
	}

	dbGorm.Ping()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start((":8080")))
}
