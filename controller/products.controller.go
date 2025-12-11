package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"echo-products-api/config"
	"echo-products-api/model"
)

func createProduct(c echo.Context) error {
	product := new(model.Products)
	db := config.DB()

	//Binding Data
	if err := c.Bind(product); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	//Create Product
	products := &model.Products{
		Name:        product.Name,
		Price:       product.Price,
		Ratings:     product.Ratings,
		Description: product.Description,
		Category:    product.Category,
	}

	if err := db.Create(&products).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": product,
	}

	return c.JSON(http.StatusOK, response)
}

func updateProduct(c echo.Context) error {
	id := c.Param("id")
	product := new(model.Products)
	db := config.DB()

	//Binding Data
	if err := c.Bind(product); err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	existingProduct := new(model.Products)

	if err := db.First(&existingProduct, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	existingProduct.Name = product.Name
	existingProduct.Price = product.Price
	existingProduct.Ratings = product.Ratings
	existingProduct.Description = product.Description
	existingProduct.Category = product.Category

	if err := db.Save(&existingProduct).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": existingProduct,
	}

	return c.JSON(http.StatusOK, response)
}

func deleteProduct(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()

	product := new(model.Products)

	err := db.Delete(&product, id).Error
	if err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"message": "Product deleted successfully",
	}

	return c.JSON(http.StatusOK, response)
}

func getProducts(c echo.Context) error {
	db := config.DB()
	products := []model.Products{}

	if err := db.Find(&products).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusInternalServerError, data)
	}

	response := map[string]interface{}{
		"data": products,
	}

	return c.JSON(http.StatusOK, response)
}

func getProductByID(c echo.Context) error {
	id := c.Param("id")
	db := config.DB()
	product := new(model.Products)

	if err := db.First(&product, id).Error; err != nil {
		data := map[string]interface{}{
			"message": err.Error(),
		}
		return c.JSON(http.StatusNotFound, data)
	}

	response := map[string]interface{}{
		"data": product,
	}
	return c.JSON(http.StatusOK, response)
}
