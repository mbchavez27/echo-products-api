package main

import {
	"net/http"
	
	"github.com/labstack/echo/v4"

	"echo-products-api/config"
	"echo-products-api/model"
}

func createProduct(c echo.Context) error {
	product := new(model.Products)
	db := config.DB()

	//Binding Data
	if err := c.Bind(product); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)}

	//Create Product
	response := map[string]interface{}{
		"data": product,
	}

	return c.JSON(http.StatusOK, response)
}