package main

import (
	"echo-products-api/config"
	"echo-products-api/controller"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//Enable CORS
	e.Use(middleware.CORS())

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
