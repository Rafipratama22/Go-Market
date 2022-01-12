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