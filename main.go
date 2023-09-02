package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DB *gorm.DB = InitDb()

func main() {
	DB.AutoMigrate(&Product{})
	r := gin.Default()
	initRoutes(r)
	r.Run()
}
