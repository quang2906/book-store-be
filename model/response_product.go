package model

type ResponseProduct struct {
	TotalPage    int       `json:"total_page"`
	TotalProduct int       `json:"total_product"`
	PageIndex    int       `json:"page_index"`
	Products     []Product `json:"products"`
}
