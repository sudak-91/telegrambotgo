package inline

type InputInvoiceMessageContent struct {
	Title                     string        `josn:"title"`
	Description               string        `json:"description"`
	Payload                   string        `json:"payload"`
	ProviderToken             string        `json:"provider_token"`
	Currency                  string        `json:"currency"`
	Prices                    []interface{} `json:"prices"` //TODO:Add LabeledPrice
	MaxTipAmount              int32         `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int32       `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string        `json:"provider_data,omitempty"`
	PhotoUrl                  string        `json:"photo_url,omitempty"`
	PhotoSize                 int32         `json:"photo_size,omitempty"`
	PhotoWidth                int32         `json:photo_width,omitempty"`
	NeedName                  bool          `json:"need_name,omitempty"`
	NeedPhoneNumber           bool          `json:"need_phone_number,omitempty"`
	NeedEmail                 bool          `json:"need_email,omitempty"`
	NeedShippingAddress       bool          `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool          `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool          `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool          `json:"is_flexible,omitempty"`
}
