package model

type Product struct {
	Id          int64   `json:"id" gorm:"primaryKey"`
	Name        string  `json:"name"`
	CategoryId  int64   `json:"category_id"`
	Image       []Image `json:"image" gorm:"foreignKey:ProductId;constraint:OnDelete:CASCADE"`
	Description string  `json:"des"`
	IsSale      bool    `json:"is_sale"`
	Price       float64 `json:"price"`
	PriceSale   float64 `json:"price_sale"`
	CreatedAt   int64   `json:"createdAt" gorm:"autoUpdateTime:milli"`
	ModifiedAt  int64   `json:"modifiedAt" gorm:"autoUpdateTime:milli"`
}
