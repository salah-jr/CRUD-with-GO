package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func InitDb() *gorm.DB {
	dsn := "root:@/crud?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting database")
	}

	return Db
}
