package types

type TelegramInlineKeyboardButton struct {
	Text                         string               `json:"text"`
	Url                          string               `json:"url,omitempty"`
	LoginUrl                     TelegramLoginUrl     `json:"login_url,omitempty"`
	CallbackData                 string               `json:"callback_data,omitempty"`
	SwitchInlineQuery            string               `json:"switch_inline_query,omitempty"`
	SwitchInlinequeryCurrentChat string               `json:"switch_inline_query_current_chat,omitempty"`
	CallbackGame                 TelegramCallbackGame `json:"callback_game,omitempty"`
	Pay                          bool                 `json:"pay,omitempty"`
}
