package inline

import types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"

type InlineQueryResultCachedVideo struct {
	Type                string                              `json:"type"`
	ID                  string                              `json:"id"`
	VideoFileID         string                              `json:"video_file_id"`
	Title               string                              `json:"title"`
	Description         string                              `json:"description,omitempty"`
	Caption             string                              `json:"caption,omitempty"`
	ParseMode           string                              `json:"parse_mode,omitempty"`
	CaptionEntities     []*types.TelegramMessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *types.TelegramInlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent                `json:"input_message_content,omitempty"`
}
