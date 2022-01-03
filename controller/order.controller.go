package controller

import (
	"github.com/Rafipratama22/go_market/service"
	"github.com/gin-gonic/gin"
	"github.com/Rafipratama22/go_market/entity"
)

type OrderController interface {
	CreateOrder(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	FindOne(ctx *gin.Context)
	UpdateOne(ctx *gin.Context)
	DeleteOne(ctx *gin.Context)
}

type orderController struct {
	service service.OrderService
}

func NewOrderController(service service.OrderService) OrderController {
	return &orderController{
		service: service,
	}
}

func (c *orderController) CreateOrder(ctx *gin.Context) {
	var body entity.Order
	ctx.ShouldBindJSON(&body)
	result := c.service.CreateOrder(body)
	ctx.JSON(200, result)
}

func (c *orderController) FindAll(ctx *gin.Context) {
	result := c.service.FindAll()
	ctx.JSON(200, result)
}

func (c *orderController) FindOne(ctx *gin.Context) {
	order_id := ctx.Param("order_id")
	result := c.service.FindOne(order_id)
	ctx.JSON(200, result)
}

func (c *orderController) UpdateOne(ctx *gin.Context) {
	order_id := ctx.Param("order_id")
	var body entity.Order
	ctx.ShouldBindJSON(&body)
	result := c.service.UpdateOne(order_id, body)
	ctx.JSON(200, result)
}

func (c *orderController) DeleteOne(ctx *gin.Context) {
	order_id := ctx.Param("order_id")
	c.service.DeleteOne(order_id)
	ctx.JSON(200, gin.H{
		"message": "success",
	})
}