package main

import "github.com/gin-gonic/gin"

func initRoutes(r *gin.Engine) {
	r.GET("/products", Products)
	r.GET("/products/:id", Show)
	r.POST("/products", Store)
	r.PATCH("/products/:id", Update)
	r.DELETE("/products/:id", Delete)
}
