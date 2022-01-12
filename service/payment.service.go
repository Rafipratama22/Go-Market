package service

import (
	// "github.com/Rafipratama22/go_market/entity"
	"fmt"

	"github.com/jinzhu/gorm"

	// "github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
)

type PaymentService interface {
	CreatePayment(payment MidtransConfig) *coreapi.ChargeResponse
}

type paymentService struct {
	db *gorm.DB
}


type MidtransConfig struct {
	OrderId string `json:"order_id"`
	TotalPrice int	`json:"total_price"`
	VaNumber string	 `json:"va_number"`
	Type int `json:"type"`
}

func NewPaymentService(db *gorm.DB) PaymentService {
	return &paymentService{
		db: db,
	}
}

func bniTransaction(va_number string, order_id string, total_price int) *coreapi.ChargeReq {
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: order_id,
			GrossAmt: int64(total_price),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.BankBni,
			VaNumber: va_number,

		},
	}
	return chargeReq
}

func mandiriTransaction(va_number string, order_id string, total_price int) *coreapi.ChargeReq {
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeEChannel,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: order_id,
			GrossAmt: int64(total_price),
		},
		EChannel: &coreapi.EChannelDetail{
			BillInfo1: "Payment for order:",
			BillInfo2: "Go Market",
		},
	}
	return chargeReq
}

func briTransaction(va_number string, order_id string, total_price int) *coreapi.ChargeReq {
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: order_id,
			GrossAmt: int64(total_price),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.BankBri,
			VaNumber: va_number,
		},
	}
	return chargeReq
}

func (c *paymentService) CreatePayment(payment MidtransConfig) *coreapi.ChargeResponse{
	var chargeReq *coreapi.ChargeReq
	if payment.Type == 1 {
		chargeReq = bniTransaction(payment.VaNumber, payment.OrderId, payment.TotalPrice)
	} else if payment.Type == 2 {
		chargeReq = mandiriTransaction(payment.VaNumber, payment.OrderId, payment.TotalPrice)
	} else if payment.Type == 3 {
		chargeReq = briTransaction(payment.VaNumber, payment.OrderId, payment.TotalPrice)
	}
	coreApi, _ := coreapi.ChargeTransaction(chargeReq)
	fmt.Println(coreApi)
	return coreApi
}
