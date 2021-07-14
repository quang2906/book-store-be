package model

type Image struct {
	ID        int64 `json:"id"`
	ProductId int64  `json:"product_id"`
	Url       string `json:"url"`
}
