package types

type PreCheckoutQuery struct {
	Id               string     `json:"id"`
	From             *User      `json:"from"`
	Currency         string     `json:"currency"`
	TotalAmount      int32      `json:"total_amount"`
	InvoicePayload   string     `json:"invoice_payload"`
	ShippingOptionId string     `json:"shipping_option_id,omitempty"`
	OrderInfo        *OrderInfo `json:"order_info"`
}
