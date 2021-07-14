package database

import (
	"log"

	"github.com/quang2906/book_store_be/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "root:tri123@tcp(127.0.0.1:3305)/shopbansach?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}
	DB = db
	db.AutoMigrate(&model.Product{}, &model.Order{}, &model.User{}, &model.OrderItem{}, &model.Category{}, &model.Image{})
}
