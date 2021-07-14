package repository

import (
	"errors"
	"time"

	"github.com/quang2906/book_store_be/database"
	"github.com/quang2906/book_store_be/model"
)

var products []*model.Product

func CreateNewProduct(product *model.Product) int64 {
	product.CreatedAt = time.Now().Unix()
	product.ModifiedAt = time.Now().Unix()
	database.DB.Create(&product)
	return product.Id
}

func GetAllProducts() []*model.Product {
	database.DB.Preload("Image").Find(&products)
	return products
}

func GetProductById(Id int64) (*model.Product, error) {
	product := new(model.Product)
	database.DB.Preload("Image").Where("id = ?", Id).Find(&product)
	if product != nil {
		return product, nil
	}
	return nil, errors.New("product not found")
}

func DeleteProductById(Id int64) error {
	rs := database.DB.Exec("delete from products where id=? ",Id)
	if rs.RowsAffected < 1 {
		return errors.New("product not found")
	}
	// DELETE from emails where id = 10 AND name = "cjinzhu"
	return nil
}

func UpdateProductById(Id int64, productUpdate *model.Product) error {
	product := new(model.Product)
	database.DB.Where("id = ?", Id).Find(&product)
	if product != nil {
		product = productUpdate
		database.DB.Save(&product)
		return nil
	}
	return errors.New("product not found")
}
