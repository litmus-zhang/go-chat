package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"log"
)

var DB *gorm.DB

func Connect(){
	dsn := "root:root@tcp(127.0.0.1:3306)/gochat?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	  log.Printf("error connecting to database %+v\n", err)
	}
}

func AutoMigrate(){

DB.AutoMigrate(models.User{})
}