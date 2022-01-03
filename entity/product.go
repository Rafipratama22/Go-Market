package entity

type Product struct {
	ID            uint64 `gorm:"primary_key;auto_increment" json:"id,omitempty"`
	Name          string `json:"name" binding:"required" gorm:"type:varchar(400)"`
	CodeNumber    string `json:"code_number" binding:"required" gorm:"type:varchar(400);unique"`
	BarcodeNumber string `json:"barcode_number" binding:"required" gorm:"type:varchar(400);unique"`
	Title         string `json:"title" binding:"required" gorm:"type:varchar(400)"`
	Description   string `json:"description" binding:"required" gorm:"type:varchar(400)"`
	Price         uint16 `json:"price" binding:"required" gorm:"type:int"`
	Stock         uint16 `json:"stock" binding:"required" gorm:"type:int"`
}
