package inline

import types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"

type InlineQueryResultGame struct {
	Type          string                              `json:"type"`
	ID            string                              `json:"id"`
	GameShortName string                              `json:"game_short_name"`
	ReplyMarkup   *types.TelegramInlineKeyboardMarkup `json:"reply_markup"`
}
