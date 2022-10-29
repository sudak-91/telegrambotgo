package types

type TelegramKeyboardButton struct {
	Text            string                          `json:"text"`
	RequestContact  bool                            `json:"request_contact,omitempty"`
	RequestLocation bool                            `json:"request_location,omitempty"`
	RequestPoll     *TelegramKeyboardButtonPollType `json:"request_poll,omitempty"`
	WebApp          *TelegramWebAppInfo             `json:"web_app"`
}
