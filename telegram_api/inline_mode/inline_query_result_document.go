package inline

import types "github.com/sudak-91/telegrambotgo/telegram_api/types"

type InlineQueryResultDocument struct {
	Type                string                      `json:"type"`
	ID                  string                      `json:"id"`
	Title               string                      `json:"title"`
	Caption             string                      `json:"caption,omitempty"`
	ParseMode           string                      `json:"parse_mode,omitempty"`
	CaptionEntities     []*types.MessageEntity      `json:"caption_entities,omitempty"`
	DocumentUrl         string                      `json:"document_url"`
	MimeType            string                      `json:"mime_type"`
	Description         string                      `json:"descritpion,omitempty"`
	ReplyMarkup         *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent        `json:"input_message_conten,omitempty"`
	ThumbUrl            string                      `json:"thumb_url,omitempty"`
	ThumbWidth          int32                       `json:"thumb_width,omitempty"`
	ThumbHeight         int32                       `json:"thumb_height,omitempty"`
}
