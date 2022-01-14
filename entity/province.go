package entity

import "time"

type Province struct {
	ID int `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Name string `gorm:"type:varchar(100);" binding:"required" json:"name"`
	ProvinceRJOID int `gorm:"int" json:"province_rjo_id"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt"`
}