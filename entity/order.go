package entity

import "time"

type Order struct {
	ID         int       `json:"id,omitempty" gorm:"auto_increment;primary_key"`
	OrderID    string    `json:"order_id" gorm:"type:string;"`
	UserID     string    `json:"user_id" gorm:"type:string"`
	AddressId  int       `json:"address_id" gorm:"type:int"`
	Status     int       `json:"status" gorm:"type:int"`
	TotalPrice int       `json:"total_price" gorm:"type:int"`
	Notes      string    `json:"notes" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamp" json:"updated_at"`
}

// PaymentTransaction []PaymentTransaction
// OrderDetail []OrderDetail `gorm:"foreignKey:order_id;"`
// PaymentTransactionID int `json:"payment_transaction_id" gorm:"type:int"`
// OrderDetailId int `json:"order_detail_id" gorm:"type:int"`
