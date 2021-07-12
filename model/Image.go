package model

type Image struct {
	ID        string `json:"id" gorm:"primaryKey"`
	ProductId int64  `json:"product_id"`
	Url       string `json:"url"`
}
