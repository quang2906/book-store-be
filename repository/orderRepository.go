package repository

import (
	"errors"
	"fmt"
	"time"

	"github.com/quang2906/book_store_be/database"
	"github.com/quang2906/book_store_be/model"
)

var orders []*model.Order

func CreateNewOrder(order *model.Order) int64 {
	order.CreatedAt = time.Now().Unix()
	order.ModifiedAt = time.Now().Unix()
	database.DB.Create(&order)
	return order.Id
}

func GetAllOrder() []*model.Order {
	database.DB.Preload("OrderItem").Find(&orders)
	return orders
}

func GetOrderById(Id int64) (*model.Order, error) {
	order := new(model.Order)
	database.DB.Preload("OrderItem").Where("id = ?", Id).Find(&order)
	if order != nil {
		return order, nil
	}
	return nil, errors.New("order not found")
}

func DeleteOrderById(Id int64) error {
	rs := database.DB.Exec("delete from orders where id=? ", Id)
	fmt.Print(rs)
	if rs.RowsAffected < 1 {
		fmt.Println(rs.RowsAffected)
		return errors.New("order not found")
	}
	// DELETE from emails where id = 10 AND name = "cjinzhu"
	return nil
}

func UpdateOrderById(Id int64, orderUpdate *model.Order) error {
	order := new(model.Order)
	database.DB.Where("id = ?", Id).Find(&order)
	if order != nil {
		order = orderUpdate
		database.DB.Save(&order)
		return nil
	}
	return errors.New("order not found")
}
