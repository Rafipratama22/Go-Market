package payment

var Mandiri = map[string]interface{}{
	"payment_type": "echannel",
    "transaction_details": map[string]interface{}{
        "order_id": nil,
        "gross_amount": nil,
    },
    "echannel" : map[string]interface{}{
        "bill_info1" : "Payment For:",
        "bill_info2" : "debt",
        "bill_info3" : "Name:",
        "bill_info4" : nil,
    },
}

type VANumber struct {
    Bank string `json:"bank"`
    VaNumber string `json:"va_number"`
}

type PaymenAmounts struct {
    PaidAt string `json:"paid_at"`
    Amount string `json:"amount"`
}

type BNIResponse struct {
    StatusCode string `json:"status_code"`
    OrderID string `json:"order_id"`
    TransactionTime string `json:"transaction_time"`
    TransactionID string `json:"transaction_id"`
    SignatureKey string `json:"signature_key"`
    GrossAmt string `json:"gross_amount"`
    PaymentType string `json:"payment_type"`
    TransactionStatus string `json:"transaction_status"`
    FraudStatus string `json:"fraud_status"`
    StatusMessage string `json:"status_message"`
    VANumber []VANumber `json:"va_number"`
}

type MandiriResponse struct {
    StatusCode string `json:"status_code"`
    StatusMessage string `json:"status_message"`
    TransactionID string `json:"transaction_id"`
    OrderID string `json:"order_id"`
    GrossAmt string `json:"gross_amount"`
    PaymentType string `json:"payment_type"`
    TransactionTime string `json:"transaction_time"`
    TransactionStatus string `json:"transaction_status"`
    ApprovalCode string `json:"approval_code"`
    SignatureKey string `json:"signature_key"`
    BillKey string `json:"bill_key"`
    BillerCode string `json:"biller_code"`
}

type PermataResponse struct {
    StatusCode string `json:"status_code"`
    StatusMessage string `json:"status_message"`
    TransactionID string `json:"transaction_id"`
    OrderID string `json:"order_id"`
    GroussAmt string `json:"gross_amount"`
    PaymentType string `json:"payment_type"`
    TransactionTime string `json:"transaction_time"`
    TransactionStatus string `json:"transaction_status"`
    FraudStatus string `json:"fraud_status"`
    PermataVaNumber string `json:"permata_va_number"`
    SignatureKey string `json:"signature_key"`
}

type BCAResponse struct {
    VANumber []VANumber `json:"va_number"`
    TransactionTime string `json:"transaction_time"`
    GrossAmt string `json:"gross_amount"`
    OrderID string `json:"order_id"`
    PaymentType string `json:"payment_type"`
    SignatureKey string `json:"signature_key"`
    StatusCode string `json:"status_code"`
    TransactionStatus string `json:"transaction_status"`
    FraudStatus string `json:"fraud_status"`
    StatusMessage string `json:"status_message"`
}

type BRIResponse struct {
    StatusCode string `json:"status_code"`
    StatusMessage string `json:"status_message"`
    TransactionID string `json:"transaction_id"`
    OrderID string `json:"order_id"`
    GrossAmt string `json:"gross_amount"`
    PaymentType string `json:"payment_type"`
    TransactionTime string `json:"transaction_time"`
    TransactionStatus string `json:"transaction_status"`
    VANumber []VANumber `json:"va_number"`
    FraudStatus string `json:"fraud_status"`
    Currency string `json:"currency"`
}