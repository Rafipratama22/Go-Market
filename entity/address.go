package entity

import "time"

type Address struct {
	ID int `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	UserID string `gorm:"user_id" json:"user_id"`
	Name string `gorm:"name" json:"name"`
	AddressLine string `gorm:"type:varchar(400);" binding:"required" json:"address_line"`
	ProvinceID int `gorm:"int REFERENCES provinces(id)" json:"province_id"`
	CityID int `gorm:"int REFERENCES cities(id)" json:"city_id"`
	KecamatanID int `gorm:"int REFERENCES kecamatans(id)" json:"kecamatan_id"`
	KelurahanID int `gorm:"int REFERNCES kelurahans(id)" json:"kelurahan_id"`
	Latitude float64 `gorm:"type:float" json:"latitude"`
	Longitude float64 `gorm:"type:float" json:"longitude"`
	Phone string `gorm:"type:varchar(20);" binding:"required" json:"phone"`
	PostalCode string `gorm:"type:varchar(10);" json:"postal_code"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}