package main

import (
	"echo-products-api/config"
	"net/http"

	"echo-products-api/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// Initialize DB
	config.DatabaseInit()
	db := config.DB()
	sqlDB, err := db.DB()
	if err != nil {
		e.Logger.Fatal("Failed to get generic database object:", err)
	}
	sqlDB.Ping()

	// Test route
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Product routes
	productsGroup := e.Group("/products")
	productsGroup.POST("", controller.CreateProduct)
	productsGroup.PUT("/:id", controller.UpdateProduct)
	productsGroup.DELETE("/:id", controller.DeleteProduct)
	productsGroup.GET("", controller.GetProducts)
	productsGroup.GET("/:id", controller.GetProductByID)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
