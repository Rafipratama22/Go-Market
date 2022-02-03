package dto

type CallbackPaymentBNI struct {
	VANumber []VANumber `json:"va_number"`
	TransactionTime string `json:"transaction_time"`
	TransactionStatus string `json:"transaction_status"`
	TransactionID string `json:"transaction_id"`
	StatusMessage string `json:"status_message"`
	StatusCode string `json:"status_code"`
	SignatureKey string `json:"signature_key"`
	SettleTime string `json:"settlement_time"`
	PaymentType string `json:"payment_type"`
	PaymentAmounts []PaymentAmounts `json:"payment_amounts"`
	OrderID string `json:"order_id"`
	MerchantID string `json:"merchant_id"`
	GrossAmount string `json:"gross_amount"`
	FraudStatus string `json:"fraud_status"`
	Currency string `json:"currency"`
}