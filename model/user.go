package model

type User struct {
	Id          int64  `json:"id" gorm:"primaryKey"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Password    []byte `json:"password"`
	Role        string `json:"role"`
	CreatedAt   int64  `json:"createdAt" gorm:"autoUpdateTime:milli"`
	ModifiedAt  int64  `json:"modifiedAt" gorm:"autoUpdateTime:milli"`
}
