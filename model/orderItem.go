package model

type OrderItem struct {
	Id        int64 `json:"id" gorm:"primaryKey"`
	ProductId int64 `json:"product_id"`
	OrderId   int64 `json:"order_id"`
	Quantity  int64 `json:"quantity"`
}
