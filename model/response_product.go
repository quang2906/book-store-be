package model

type ResponseProduct struct {
	TotalPage    int       `json:"total_page"`
	TotalProduct int       `json:"total_product"`
	Products     []Product `json:"products"`
}