package entity

type Address struct {
	ID int `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Address_line string `gorm:"type:varchar(400);" binding:"required" json:"address_line"`
	Province_id int `gorm:"int REFERENCES provinces(id)" json:"province_id"`
	City_id int `gorm:"int REFERENCES cities(id)" json:"city_id"`
	Kecamatan_id int `gorm:"int REFERENCES kecamatans(id)" json:"kecamatan_id"`
	Kelurahan_id int `gorm:"int REFERNCES kelurahans(id)"`
}