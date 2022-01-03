package service

import (
	"github.com/Rafipratama22/go_market/repository"
	"github.com/Rafipratama22/go_market/entity"
)


type OrderService interface {
	CreateOrder(body entity.Order) entity.Order
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

func (c *orderService) CreateOrder(body entity.Order) entity.Order {
	return c.repository.CreateOrder(body)
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