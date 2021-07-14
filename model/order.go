package model

type Order struct {
	Id           int64       `json:"id" gorm:"primaryKey"`
	OrderItem    []OrderItem `json:"order_item" gorm:"foreignKey:order_id;constraint:OnDelete:CASCADE;constraint:OnUpdate:CASCADE"`
	TotalProduct int         `json:"total_product"`
	TotalAmount  float64     `json:"total_amount"`
	Customer     string      `json:"customer"`
	Address      string      `json:"address"`
	Phone        string      `json:"phone"`
	Email        string      `json:"email"`
	Payment      string      `json:"payment"`
	CreatedAt    int64       `json:"createdAt" gorm:"autoUpdateTime:milli"`
	ModifiedAt   int64       `json:"modifiedAt" gorm:"autoUpdateTime:milli"`
}
