package entity

import (
	// "github.com/midtrans/midtrans-go/coreapi"
)

type PaymentTransaction struct {
	ID int `json:"id,omitempty" gorm:"auto_increment;primary_key"`
	TransactionID string `json:"transaction_id" gorm:"type:varchar(100);primary_key"`
	OrderID string `json:"order_id" gorm:"type:varchar(100)"`
	PaymentCredentialID string `json:"payment_credential_id" gorm:"type:varchar(100)"`
	TransactionStatus string `json:"transaction_status" gorm:"type:varchar(100)"`
	Information []byte `json:"information" gorm:"type:json"`
	Type int `json:"type" gorm:"type:int"`
	TotalPrice int `json:"total_price" gorm:"type:int"`
}