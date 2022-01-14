package repository

import (
	// "fmt"

	"github.com/Rafipratama22/go_market/entity"
	// "github.com/jinzhu/gorm"
	"gorm.io/gorm"
)

type OrderRepository interface {
	CreateOrder(body entity.Order, order_detail []entity.OrderDetail, address entity.Address, boolAddress bool) (entity.Order)
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

func (c *orderRepository) CreateOrder(body entity.Order, order_detail []entity.OrderDetail, address entity.Address, boolAddress bool) entity.Order{
	c.db.Model(&order_detail).Create(&order_detail)
	// fmt.Println(order_details)
	if !boolAddress {
		c.db.Model(&address).Create(&address)
	}
	// fmt.Println(addresses)
	body.AddressId = address.ID
	body.Status = 1
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
	c.db.Model(&result).Where("order_id = ?", order_id).Updates(&result)
	return result
}

func (c *orderRepository) DeleteOne(order_id string) {
	var order entity.Order
	c.db.Where("order_id = ?", order_id).First(&order)
	c.db.Delete(&order)
}