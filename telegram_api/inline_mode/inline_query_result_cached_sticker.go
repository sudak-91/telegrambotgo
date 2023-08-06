package inline

import types "github.com/sudak-91/telegrambotgo/telegram_api/types"

type InlineQueryResultCachedSticker struct {
	Type                string                      `json:"type"`
	ID                  string                      `json:"id"`
	StickerFileID       string                      `json:"sticker_file_id"`
	ReplyMarkup         *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent        `json:"input_message_content,omitempty"`
}
