package inline

import types "github.com/sudak-91/telegrambotgo/telegram_api/types"

type InlineQueryResultPhoto struct {
	Type                string                      `json:"type"`
	ID                  string                      `json:"id"`
	PhotoURL            string                      `json:"photo_url"`
	ThumbURL            string                      `json:"thumb_url"`
	PhotoWidth          int32                       `json:"photo_width,omitempty"`
	PhotoHeight         int32                       `json:"photo_height,omitempty"`
	Title               string                      `json:"title,omitempty"`
	Description         string                      `json:"description,omitempty"`
	Caption             string                      `json:"caption,omitempty"`
	ParseMode           string                      `json:"parse_mode,omitempty"`
	CaptionEntities     []*types.MessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent        `json:"input_message_content,omitempty"`
}
