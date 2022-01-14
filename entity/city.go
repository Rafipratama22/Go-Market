package entity

import "time"

type City struct {
	ID int `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Name string `gorm:"type:varchar(100);" binding:"required" json:"name"`
	Province int `gorm:"int REFERENCES provinces(id)" json:"area_province_id"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt"`
	CityCode int `gorm:"type:int" json:"city_code"`
	CityRJOID int `gorm:"type:int" json:"city_rjo_id"`
}