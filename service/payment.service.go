package service

import (
	"github.com/Rafipratama22/go_market/entity"
	// "github.com/Rafipratama22/go_market/dto"
	"fmt"
	"gorm.io/gorm"
	"encoding/json"
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

func bcaTransaction(va_number string, order_id string, total_price int) *coreapi.ChargeReq {
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: order_id,
			GrossAmt: int64(total_price),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.BankBca,
			VaNumber: va_number,
			Bca: &coreapi.BcaBankTransferDetail{
				SubCompanyCode: "00000",
			},
		},
	}
	return chargeReq
}

func permataTransaction(va_number string, order_id string, total_price int) *coreapi.ChargeReq {
	chargeReq := &coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeBankTransfer,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: order_id,
			GrossAmt: int64(total_price),
		},
		BankTransfer: &coreapi.BankTransferDetails{
			Bank: midtrans.BankPermata,
			VaNumber: va_number,
			Permata: &coreapi.PermataBankTransferDetail{
				RecipientName: "Rafipratama",
			},
		},
	}
	return chargeReq
}

func creditCardTransaction() {
	// chargeReq := 
}

func (c *paymentService) CreatePayment(payment MidtransConfig) *coreapi.ChargeResponse{
	var chargeReq *coreapi.ChargeReq
	var resultCharge entity.PaymentTransaction
	if payment.Type == 1 {
		chargeReq = bniTransaction(payment.VaNumber, payment.OrderId, payment.TotalPrice)
	} else if payment.Type == 2 {
		chargeReq = mandiriTransaction(payment.VaNumber, payment.OrderId, payment.TotalPrice)
	} else if payment.Type == 3 {
		chargeReq = briTransaction(payment.VaNumber, payment.OrderId, payment.TotalPrice)
	} else if payment.Type == 4 {
		chargeReq = bcaTransaction(payment.VaNumber, payment.OrderId, payment.TotalPrice)
	} else if payment.Type == 5 {
		chargeReq = permataTransaction(payment.VaNumber, payment.OrderId, payment.TotalPrice)
	}
	coreApi, _ := coreapi.ChargeTransaction(chargeReq)
	fmt.Println(coreApi)
	resultCharge.OrderID = payment.OrderId
	resultCharge.TransactionID = coreApi.TransactionID
	resultCharge.TransactionStatus = coreApi.TransactionStatus
	resultCharge.PaymentCredentialID = ""
	resultCharge.Type = payment.Type
	changeJson, _ := json.Marshal(coreApi)
	resultCharge.Information = changeJson
	resultCharge.TotalPrice = payment.TotalPrice
	c.db.Model(&resultCharge).Create(&resultCharge)
	c.db.Model(&entity.Order{}).Where("order_id = ?", payment.OrderId).Update("payment_status", 1)
	fmt.Println(resultCharge)
	return coreApi
}

func (c *paymentService) CallbackPayment() {
	// var order entity.Order
}