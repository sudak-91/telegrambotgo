package inline

import types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"

type InlineQueryResultCachedDocument struct {
	Type                string                              `json:"type"`
	ID                  string                              `json:"id"`
	Title               string                              `json:"title"`
	DocumentFileID      string                              `json:"document_file_id"`
	Description         string                              `json:"description,omitempty"`
	Caption             string                              `json:"caption,omitempty"`
	ParseMode           string                              `json:"parse_mode,omitempty"`
	CaptionEntities     []*types.TelegramMessageEntity      `json:"caption_entities,omitempty"`
	ReplyMarkup         *types.TelegramInlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent                `json:"input_message_content,omitempty"`
}
