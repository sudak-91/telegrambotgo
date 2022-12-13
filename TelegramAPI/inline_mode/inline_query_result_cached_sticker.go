package inline

import types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"

type InlineQueryResultCachedSticker struct {
	Type                string                              `json:"type"`
	ID                  string                              `json:"id"`
	StickerFileID       string                              `json:"sticker_file_id"`
	ReplyMarkup         *types.TelegramInlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent                `json:"input_message_content,omitempty"`
}
