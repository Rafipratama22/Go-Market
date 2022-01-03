package entity

type OrderDetail struct {
	ID int `json:"id,omitempty" gorm:"auto_increment;primary_key"`
	ProductId int `json:"product_id" binding:"required" gorm:"type:int"`
	UserId int `json:"user_id" binding:"required" gorm:"type:int"`
	Note string `json:"note" binding:"required" gorm:"type:varchar(400)"`
	Quantity int `json:"quantity" binding:"required" gorm:"type:int"`
	OrderId int `json:"order_id" binding:"required" gorm:"type:int"`
	Price int `json:"price" binding:"required" gorm:"type:int"`
}