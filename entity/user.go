package entity

type User struct {
	ID       int `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Name     string `json:"name"  gorm:"type:varchar(400)"`
	Email    string `json:"email"  gorm:"type:varchar(400);unique"`
	Password string `json:"password"  gorm:"type:varchar(400)"`
	Phone    string `json:"phone"  gorm:"type:varchar(400)"`
	Address  string `json:"address"  gorm:"type:varchar(400)"`
	City     string `json:"city"  gorm:"type:varchar(400)"`
	Province string `json:"province"  gorm:"type:varchar(400)"`
	Country  string `json:"country"  gorm:"type:varchar(400)"`
	ZipCode  string `json:"zip_code"  gorm:"type:varchar(400)"`
	Role     string `json:"role"  gorm:"type:varchar(400)"`

}
