package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/Rafipratama22/go_market/service"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentController interface {
	CreatePayment(ctx *gin.Context)
}

type paymentController struct {
	paymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) PaymentController {
	return &paymentController{
		paymentService: paymentService,
	}
}

func (c *paymentController) CreatePayment(ctx *gin.Context) {
	var payment service.MidtransConfig
	var returnCharge *coreapi.ChargeResponse
	err := ctx.ShouldBindJSON(&payment)
	if err != nil {
		panic(err)
	}
	returnCharge = c.paymentService.CreatePayment(payment)
	ctx.JSON(200, returnCharge)
}