package service

import (
	"fmt"

	"github.com/Rafipratama22/go_market/dto"
	"github.com/Rafipratama22/go_market/entity"
	"github.com/Rafipratama22/go_market/repository"
	"github.com/google/uuid"
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

func (c *orderService) CreateOrder(body dto.CreateOrder) (entity.Order) {
	var order entity.Order
	// var order_detail entity.OrderDetail
	order_id := uuid.New().String()
	// var location []byte
	// fmt.Println("order_id => " + order_id)
	address := entity.Address{
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
	}
	var	order_detail []entity.OrderDetail
	var total_price int
	for i := 0; i < len(body.OrderDetail); i++ {
		order_details := entity.OrderDetail{
			OrderID: order_id,
			ProductBarcode: body.OrderDetail[i].ProductBarcode,
			Price: body.OrderDetail[i].Price,
			Quantity: body.OrderDetail[i].Quantity,
		}
		total_price += body.OrderDetail[i].Price * body.OrderDetail[i].Quantity
		fmt.Println(order_details)
		order_detail = append(order_detail, order_details)
	}
	order = entity.Order{
		OrderID: order_id,
		UserID: body.UserID,
		Notes: body.Notes,
	}
	// fmt.Println(address)
	c.repository.CreateOrder(order, order_detail, address)
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