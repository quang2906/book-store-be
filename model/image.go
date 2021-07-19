package model

type Image struct {
	Id        int64  `json:"id" gorm:"primaryKey"`
	ProductId int64  `json:"product_id"`
	Url       string `json:"url"`
}
