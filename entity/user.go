package entity

type User struct {
	ID       int `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Name     string `json:"name" binding:"required" gorm:"type:varchar(400)"`
	Email    string `json:"email" binding:"required" gorm:"type:varchar(400);unique"`
	Password string `json:"password" binding:"required" gorm:"type:varchar(400)"`
	Phone    string `json:"phone" binding:"required" gorm:"type:varchar(400)"`
	Address  string `json:"address" binding:"required" gorm:"type:varchar(400)"`
	City     string `json:"city" binding:"required" gorm:"type:varchar(400)"`
	Province string `json:"province" binding:"required" gorm:"type:varchar(400)"`
	Country  string `json:"country" binding:"required" gorm:"type:varchar(400)"`
	ZipCode  string `json:"zip_code" binding:"required" gorm:"type:varchar(400)"`
	Role     string `json:"role" binding:"required" gorm:"type:varchar(400)"`

}
