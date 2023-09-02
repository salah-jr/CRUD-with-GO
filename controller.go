package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Products(context *gin.Context) {
	var products []Product
	DB.Find(&products)

	context.JSON(http.StatusOK, gin.H{
		"data":    products,
		"message": "All Products",
	})
}

func Show(context *gin.Context) {
	product := getByID(context)

	if product.ID == 0 {
		return
	}
	context.JSON(http.StatusFound, gin.H{
		"data":    product,
		"message": "",
	})
}

func Store(context *gin.Context) {
	var product Product

	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	DB.Create(&product)

	context.JSON(http.StatusCreated, gin.H{
		"data":    product,
		"message": "Product Created!",
	})
}

func Update(context *gin.Context) {
	product := getByID(context)

	if product.ID == 0 {
		return
	}

	if err := context.ShouldBindJSON(&product); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	DB.Save(&product)

	context.JSON(http.StatusCreated, gin.H{
		"data":    product,
		"message": "Product updated",
	})
}

func Delete(context *gin.Context) {
	product := getByID(context)

	if product.ID == 0 {
		return
	}

	// sot delete
	//Db.Delete(&product)

	// Delete permanently
	DB.Unscoped().Delete(&product)

	context.JSON(http.StatusOK, gin.H{
		"data":    "",
		"message": "Product deleted",
	})
}

func getByID(context *gin.Context) Product {
	var product Product
	id := context.Param("id")
	DB.First(&product, id)
	if product.ID == 0 {
		context.JSON(http.StatusNotFound, gin.H{
			"data":    "",
			"message": "Product doesn't exist",
		})
	}

	return product
}
