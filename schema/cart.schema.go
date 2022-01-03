package schema

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Cart struct {
	ID primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ProductID string `json:"product_id" binding:"required" gorm:"type:varchar(400)"`
	Quantity int `json:"quantity" gorm:"type:int"`
	UserID string `json:"user_id" binding:"required" gorm:"type:varchar(400)"`
	// OrderID string `json:"order_id" binding:"required"`
}