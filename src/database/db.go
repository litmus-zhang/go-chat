package database

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB
func Connect(){
	// import "gorm.io/driver/mysql"
	// refer: https://gorm.io/docs/connecting_to_the_database.html#MySQL
	dsn := "root:root@tcp(127.0.0.1:3306)/gochat?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
	  log.Printf("error connecting to database %+v\n", err)
	}
}

func AutoMigrate(){

}