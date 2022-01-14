package entity

import "time"

type Kelurahan struct {
	ID int `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Name string `gorm:"type:varchar(100);" binding:"required" json:"name"`
	Kecamatan int `gorm:"int REFERENCES kecamatans(id)" json:"area_kecamatan_id"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updatedAt"`
}