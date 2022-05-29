package types

type TelegramContact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int32  `json:"user_id,omitempty"`
	VCard       string `json:"vcard,omitempty"`
}
