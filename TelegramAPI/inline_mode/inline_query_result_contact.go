package inline

import types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"

type InlineQueryResultContact struct {
	Type                string                              `json:"type"`
	ID                  string                              `json:"id"`
	PhoneNumber         string                              `json:"phone_number"`
	FirstName           string                              `json:"first_name"`
	LastName            string                              `json:"last_name,omitempty"`
	VCard               string                              `json:"vcard,omitempty"`
	ReplyMarkup         *types.TelegramInlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent                `json:"input_message_content,omitempty"`
	ThumbUrl            string                              `json:"thumb_url,omitempty"`
	ThumbWidth          int32                               `json:"thumb_width,omitempty"`
	ThumbHeight         int32                               `json:"thumb_height,omitempty"`
}
