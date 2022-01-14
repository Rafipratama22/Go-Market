package dto

type OrderDetail struct {
	OrderID string `json:"order_id" binding:"required"`
	ProductBarcode string `json:"product_barcode" binding:"required"`
	Price int `json:"price" binding:"required"`
	Quantity int `json:"quantity" binding:"required"`
}

type CreateOrder struct {
	Address Address `json:"address" binding:"required"`
	OrderDetail []OrderDetail `json:"order_detail" binding:"required"`
	UserID string `json:"user_id" binding:"required"`
	Notes string `json:"notes" `
}

type Address struct {
	AddressLine string `json:"address_line" binding:"required"`
	ProvinceID int `json:"province_id" binding:"required"`
	CityID int `json:"city_id" binding:"required"`
	KecamatanID int `json:"kecamatan_id" binding:"required"`
	KelurahanID int `json:"kelurahan_id" binding:"required"`
	Phone string `json:"phone" binding:"required"`
	Name string `json:"name" binding:"required"`
	PostalCode string `json:"postal_code"`
	Longitude float64 `json:"longitude"`
	Latitude float64 `json:"latitude"`
}