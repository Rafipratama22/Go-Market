package controller

import (
	"github.com/Rafipratama22/go_market/schema"
	"github.com/Rafipratama22/go_market/service"
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/bson"
)

type CartController interface {
	CreateCart(ctx *gin.Context)
	FindAllCart(ctx *gin.Context)
	FindOneCart(ctx *gin.Context)
	UpdateOneCart(ctx *gin.Context)
	DeleteOneCart(ctx *gin.Context)
}

type cartController struct {
	cartService service.CartService
}

func NewCartController(service service.CartService) CartController {
	return &cartController{
		cartService: service,
	}
}

func (c *cartController) CreateCart(ctx *gin.Context) {
	var cart schema.Cart
	err := ctx.ShouldBindJSON(&cart)
	if err != nil {
		panic(err)
	}
	res := c.cartService.CreateCart(cart)
	ctx.JSON(201, res)
}

func (c *cartController) FindAllCart(ctx *gin.Context) {
	res := c.cartService.FindAll()
	ctx.JSON(200, res)
}

func (c *cartController) FindOneCart(ctx *gin.Context) {
	order_id := ctx.Param("order_id")
	cartOne := c.cartService.FindOne(order_id)
	ctx.JSON(200, cartOne)
}

func (c *cartController) UpdateOneCart(ctx *gin.Context) {
	var cart schema.Cart
	order_id := ctx.Param("order_id")
	err := ctx.ShouldBindJSON(&cart)
	if err != nil {
		panic(err)
	}
	cartUpdateOne := c.cartService.UpdateOne(order_id, cart)
	ctx.JSON(201, cartUpdateOne)
}

func (c *cartController) DeleteOneCart(ctx *gin.Context) {
	order_id := ctx.Param("order_id")
	cartDeleteOne := c.cartService.DeleteOne(order_id)
	ctx.JSON(200, cartDeleteOne)
}