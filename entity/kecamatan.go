package entity

import "time"

type Kecamatan struct {
	ID int `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Name string `gorm:"type:varchar(100);" binding:"required" json:"name"`
	City int `gorm:"int REFERENCES cities(id)" json:"area_city_id"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt"`
	
}