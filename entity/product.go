package entity

type Product struct {
	ID                  int `gorm:"primary_key;auto_increment;" json:"id,omitempty"`
	Name                string `json:"name" binding:"required" gorm:"type:text(1000);"`
	CodeNumber          string `json:"code_number" binding:"required" gorm:"type:varchar(100);unique;"`
	BarcodeNumber       string `json:"barcode_number" binding:"required" gorm:"type:varchar(100);unique;"`
	Title               string `json:"title" binding:"required" gorm:"type:varchar(400);"`
	Description         string `json:"description" binding:"required" gorm:"type:longtext"`
	Price               int    `json:"price" binding:"required" gorm:"type:int;"`
	Stock               int    `json:"stock" binding:"required" gorm:"type:int;"`
	Department          string `json:"department" gorm:"type:varchar(100);"`
	Department_id       int    `json:"department_id" gorm:"type:int;" `
	Category            string `json:"category" gorm:"type:varchar(100);"`
	Category_id         int    `json:"category_id" gorm:"type:int;"`
	Sub_Category        string `json:"sub_category" gorm:"type:varchar(100)"`
	Sub_Category_id     int    `gorm:"type:int" json:"sub_category_id"`
	Sub_sub_category    string `gorm:"type:varchar(400)" json:"sub_sub_category"`
	Sub_sub_category_id int    `gorm:"type:int" json:"sub_sub_category_id"`
	Brand               string `gorm:"type:varchar(100)" json:"brand"`
	Brand_id            int16  `gorm:"type:int" json:"brand_id"`
	Available_stock     int    `gorm:"type:int" json:"available_stock"`
	Color               string `gorm:"type:varchar(100)" json:"color"`
	Priority            int    `gorm:"type:int" json:"priority"`
	Lukot               string `gorm:"type:varchar(100)" json:"lukot"`
	Weight              int    `gorm:"type:int" json:"weight"`
	Weight_unit         string `gorm:"type:varchar(100)" json:"weight_unit"`
	Size                string `gorm:"type:varchar(100)" json:"size"`
	Size_unit           string `gorm:"type:varchar(100)" json:"size_unit"`
	Caption             string `gorm:"type:text(1000)" json:"caption"`
	Hastag              string `gorm:"type:text(1000)" json:"hastag"`
	Promo_price         int    `gorm:"type:int" json:"promo_price"`
	Created_by          int    `gorm:"type:int" json:"created_by"`
	Updated_by          int    `gorm:"type:int" json:"updated_by"`
	Is_active           bool   `gorm:"type:bool" json:"is_active"`
	Image_1_Original    string `gorm:"type:text(1000)" json:"image_1_original"`
	Image_1_Medium      string `gorm:"type:text(1000)" json:"image_1_medium"`
	Image_1_Small       string `gorm:"type:text(1000)" json:"image_1_small"`
	Image_1_Thumbnail   string `gorm:"type:text(1000)" json:"image_1_thumbnail"`
	Image_2_Original    string `gorm:"type:text(1000)" json:"image_2_original"`
	Image_2_Medium      string `gorm:"type:text(1000)" json:"image_2_medium"`
	Image_2_Small       string `gorm:"type:text(1000)" json:"image_2_small"`
	Image_2_Thumbnail   string `gorm:"type:text(1000)" json:"image_2_thumbnail"`
	Image_3_Original    string `gorm:"type:text(1000)" json:"image_3_original"`
	Image_3_Medium      string `gorm:"type:text(1000)" json:"image_3_medium"`
	Image_3_Small       string `gorm:"type:text(1000)" json:"image_3_small"`
	Image_3_Thumbnail   string `gorm:"type:text(1000)" json:"image_3_thumbnail"`
	Image_4_Original    string `gorm:"type:text(1000)" json:"image_4_original"`
	Image_4_Medium      string `gorm:"type:text(1000)" json:"image_4_medium"`
	Image_4_Small       string `gorm:"type:text(1000)" json:"image_4_small"`
	Image_4_Thumbnail   string `gorm:"type:text(1000)" json:"image_4_thumbnail"`
	Smarthchoice        string `gorm:"type:text(1000)" json:"smarthchoice"`
	Op_fresh            string `gorm:"type:varchar(100)" json:"op_fresh"`
	Dietary             string `gorm:"type:varchar(100)" json:"dietary"`
	Import              bool   `gorm:"type:bool" json:"import"`
}
