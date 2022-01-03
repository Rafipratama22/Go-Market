package entity

type Order struct {
	ID int `json:"id,omitempty" gorm:"auto_increment;primary_key"`
	UserId int `json:"user_id" gorm:"type:int"`
	OrderDetailId int `json:"order_detail_id" gorm:"type:int"`
	AddressId int `json:"address_id" gorm:"type:int"`
	Status int `json:"status" gorm:"type:int"`
	TotalPrice int `json:"total_price" gorm:"type:int"`
}