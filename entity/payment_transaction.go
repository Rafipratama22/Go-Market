package entity

import "time"
type PaymentTransaction struct {
	ID int `json:"id,omitempty" gorm:"auto_increment;primary_key"`
	TransactionID string `json:"transaction_id" gorm:"type:varchar(100)"`
	OrderID string `json:"order_id" gorm:"type:varchar(100)"`
	Order Order `gorm:"foreignKey:order_id;"`
	PaymentCredentialID string `json:"payment_credential_id" gorm:"type:varchar(100)"`
	TransactionStatus string `json:"transaction_status" gorm:"type:varchar(100)"`
	Information []byte `json:"information" gorm:"type:json"`
	Type int `json:"type" gorm:"type:int"`
	TotalPrice int `json:"total_price" gorm:"type:int"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}