package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/quang2906/book_store_be/database"
	"github.com/quang2906/book_store_be/model"
)

var (
	products []*model.Product
	limit    int = 6
)

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

func GetProductByCategory(categoryId int64, offset int) ([]*model.Product, error) {
	var products []*model.Product
	sql := "SELECT * FROM products where category_id=? limit 6 offset ?;"
	database.DB.Raw(sql, categoryId, offset).Scan(&products)
	if products != nil {
		return products, nil
	}
	return nil, errors.New("product not found")
}

func DeleteProductById(Id int64) error {
	rs := database.DB.Exec("delete from products where id=? ", Id)
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

func SearchProductRepo(nameProduct string, offset int) []model.Product {
	key := "%" + nameProduct + "%"
	var products []model.Product
	// sql := "select * from products where products.name LIKE ? limit 6 offset ?"
	// fmt.Print(sql)
	// database.DB.Preload("Image").Raw(sql, key, offset).Scan(&products)
	// fmt.Println(database.DB.Preload("Image").Raw(sql, key).Offset(offset))
	// fmt.Println(products)

	database.DB.Preload("Image").Where("products.name LIKE ?", key).Limit(limit).Offset(offset).Find(&products)
	return products
}

func SortProductRepo(sort string, offset int) []model.Product {
	var products []model.Product
	// priceAsc := "select * from products order by products.price limit 6 offset ?"
	// priceDesc := "select * from products order by products.price desc limit 6 offset ?"
	// nameAsc := "select * from products order by products.name limit 6 offset ?"
	// nameDesc := "select * from products order by products.name desc limit 6 offset ?"
	// if sort == "priceasc" {
	// 	database.DB.Raw(priceAsc, offset).Scan(&products)
	// 	fmt.Println(products)
	// }
	// if sort == "pricedesc" {
	// 	database.DB.Raw(priceDesc, offset).Scan(&products)
	// }
	// if sort == "nameasc" {
	// 	database.DB.Raw(nameAsc, offset).Scan(&products)
	// }
	// if sort == "namedesc" {
	// 	fmt.Println("1")
	// 	database.DB.Raw(nameDesc, offset).Scan(&products)
	// }

	if sort == "priceasc" {
		database.DB.Preload("Image").Order("products.price").Limit(limit).Offset(offset).Find(&products)
		fmt.Println(products)
	}
	if sort == "pricedesc" {
		database.DB.Preload("Image").Order("products.price desc").Limit(limit).Offset(offset).Find(&products)
	}
	if sort == "nameasc" {
		database.DB.Preload("Image").Order("products.name").Limit(limit).Offset(offset).Find(&products)
	}
	if sort == "namedesc" {
		database.DB.Preload("Image").Order("products.name desc").Limit(limit).Offset(offset).Find(&products)
	}
	return products
}

func TotalProduct(nameProduct string) int {
	var totalProducts int
	key := "%" + nameProduct + "%"
	database.DB.Raw("select count(*) from products where products.name LIKE ?", key).Scan(&totalProducts)
	return totalProducts
}
