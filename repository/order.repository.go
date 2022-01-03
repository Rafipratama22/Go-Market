package repository

import (
	"github.com/Rafipratama22/go_market/entity"
	"github.com/jinzhu/gorm"
)

type OrderRepository interface {
	CreateOrder(body entity.Order) entity.Order
	FindAll() []entity.Order
	FindOne(order_id string) entity.Order
	UpdateOne(order_id string, body entity.Order) entity.Order
	DeleteOne(order_id string)
}

type orderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository{
	return &orderRepository{
		db: db,
	}
}

func (c *orderRepository) CreateOrder(body entity.Order) entity.Order{
	c.db.Create(&body)
	return body
}

func (c *orderRepository) FindAll() []entity.Order {
	var result []entity.Order
	c.db.Find(&result)
	return result
}

func (c *orderRepository) FindOne(order_id string) entity.Order {
	var result entity.Order
	c.db.Where("order_id = ?", order_id).First(&result)
	return result
}

func (c *orderRepository) UpdateOne(order_id string, body entity.Order) entity.Order {
	var result entity.Order
	c.db.Model(&result).Where("order_id = ?", order_id).Update(&result)
	return result
}

func (c *orderRepository) DeleteOne(order_id string) {
	var order entity.Order
	c.db.Where("order_id = ?", order_id).First(&order)
	c.db.Delete(&order)
}