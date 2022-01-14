package entity

type OrderDetail struct {
	ID int `json:"id,omitempty" gorm:"auto_increment;primary_key"`
	ProductBarcode string`json:"product_barcode" binding:"required" gorm:"type:varchar(100)"`
	UserId string `json:"user_id" binding:"required" gorm:"type:varchar(100)"`
	Note string `json:"note" binding:"required" gorm:"type:varchar(400)"`
	Quantity int `json:"quantity" binding:"required" gorm:"type:int"`
	OrderID string `json:"order_id" binding:"required" gorm:"type:varchar(100)"`
	Price int `json:"price" binding:"required" gorm:"type:int"`
}