package service

import (
	"fmt"

	"github.com/Rafipratama22/go_market/dto"
	"github.com/Rafipratama22/go_market/entity"
	"github.com/Rafipratama22/go_market/repository"
	// "github.com/google/uuid"
	"crypto/rand"
)


type OrderService interface {
	CreateOrder(body dto.CreateOrder) (entity.Order)
	FindAll() []entity.Order
	FindOne(order_id string) entity.Order
	UpdateOne(order_id string, body entity.Order) entity.Order
	DeleteOne(order_id string)
}

type orderService struct {
	repository repository.OrderRepository
}

func NewOrderService(repository repository.OrderRepository) OrderService {
	return &orderService{
		repository: repository,
	}
}
func makeUUID() string {
	n := 5
	b := make([]byte, n)
	if _, err := rand.Read(b); err != nil {
		panic(err)
	}
	s := fmt.Sprintf("%X", b)
	fmt.Println(s)
	return s

}

func (c *orderService) CreateOrder(body dto.CreateOrder) (entity.Order) {
	var order entity.Order
	// var order_detail entity.OrderDetail
	order_id := makeUUID()
	var address entity.Address
	var boolAddress bool
	// var location []byte
	// fmt.Println("order_id => " + order_id)
	if body.Address.AddressID == 0{
		boolAddress = false
	} else {
		boolAddress = true
	}
	address = entity.Address{
		AddressLine: body.Address.AddressLine,
		ProvinceID: body.Address.ProvinceID,
		CityID: body.Address.CityID,
		KecamatanID: body.Address.KecamatanID,
		KelurahanID: body.Address.KelurahanID,
		Phone: body.Address.Phone,
		Name: body.Address.Name,
		UserID: body.UserID,
		Longitude: body.Address.Longitude,
		Latitude: body.Address.Latitude,
		PostalCode: body.Address.PostalCode,
	}
	var	order_detail []entity.OrderDetail
	var total_price int
	for i := 0; i < len(body.OrderDetail); i++ {
		order_details := entity.OrderDetail{
			OrderID: order_id,
			ProductBarcode: body.OrderDetail[i].BarcodeNumber,
			Price: body.OrderDetail[i].Price,
			Quantity: body.OrderDetail[i].Quantity,
			UserId: body.UserID,
		}
		total_price += body.OrderDetail[i].Price * body.OrderDetail[i].Quantity
		fmt.Println(total_price)
		order_detail = append(order_detail, order_details)
	}
	order = entity.Order{
		OrderID: order_id,
		UserID: body.UserID,
		Notes: body.Notes,
		TotalPrice: total_price,
		Status: 1,
	}
	// fmt.Println(address)
	c.repository.CreateOrder(order, order_detail, address, boolAddress)
	return order
}

func (c *orderService) FindAll() []entity.Order {
	return c.repository.FindAll()
}

func (c *orderService) FindOne(order_id string) entity.Order {
	return c.repository.FindOne(order_id)
}

func (c *orderService) UpdateOne(order_id string, body entity.Order) entity.Order {
	return c.repository.UpdateOne(order_id, body)
}

func (c *orderService) DeleteOne(order_id string) {
	c.repository.DeleteOne(order_id)
}


// student-03-001ce81a9a9f@qwiklabs.net
// student-03-6230df30116e@qwiklabs.net
// cAUAQZNCJkz6
// qwiklabs-gcp-03-85bc31ec833e
// qwiklabs-gcp-03-279d29713d38