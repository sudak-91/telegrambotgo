package inline

import types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"

type InlineQueryResultArticle struct {
	Type                string                              `json:"type"`
	ID                  string                              `json:"id"`
	Title               string                              `json:"title"`
	InputMessageContent *InputMessageContent                `json:"input_message_content"`
	ReplyMarkup         *types.TelegramInlineKeyboardMarkup `json:"reply_markup,omitempty"`
	URL                 string                              `json:"utl,omitempty"`
	HideUrl             bool                                `json:"hide_url,omitempty"`
	Description         string                              `json:"description,omitempty"`
	ThumbUrl            string                              `json:"thumb_url,omitempty"`
	ThumbWidth          int32                               `json:"thumb_width,omitempty"`
	ThumbHeight         int32                               `json:"thumb_height,omitempty"`
}
