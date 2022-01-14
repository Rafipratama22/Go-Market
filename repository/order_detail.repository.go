package repository

import (
	"github.com/Rafipratama22/go_market/entity"
	"gorm.io/gorm"
)

type OrderDetailRepository interface {
	CreateOrderDetail(body entity.OrderDetail) entity.OrderDetail
	FindAll() []entity.OrderDetail
	FindOne(order_id string) entity.OrderDetail
	UpdateOne(order_id string) entity.OrderDetail
	DeleteOne(order_id string)
}

type orderDetailRepository struct {
	db *gorm.DB
}

func NewOrderDetailRepository(db *gorm.DB) OrderDetailRepository {
	return &orderDetailRepository{
		db: db,
	}
}

func (c *orderDetailRepository) CreateOrderDetail(body entity.OrderDetail) entity.OrderDetail {
	c.db.Create(&body)
	return body
}

func (c *orderDetailRepository) FindAll() []entity.OrderDetail {
	var result []entity.OrderDetail
	c.db.Find(&result)
	return result
}

func (c *orderDetailRepository) FindOne(order_id string) entity.OrderDetail {
	var result entity.OrderDetail
	c.db.Where("order_id = ?", order_id).First(&result)
	return result
}

func (c *orderDetailRepository) UpdateOne(order_id string) entity.OrderDetail {
	var result entity.OrderDetail
	c.db.Model(&result).Where("order_id = ?", order_id).Updates(&result)
	return result
}

func (c *orderDetailRepository) DeleteOne(order_id string) {
	var order entity.OrderDetail
	c.db.Where("order_id = ?", order_id).First(&order)
	c.db.Delete(&order)
}