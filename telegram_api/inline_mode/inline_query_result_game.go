package inline

import types "github.com/sudak-91/telegrambotgo/telegram_api/types"

type InlineQueryResultGame struct {
	Type          string                      `json:"type"`
	ID            string                      `json:"id"`
	GameShortName string                      `json:"game_short_name"`
	ReplyMarkup   *types.InlineKeyboardMarkup `json:"reply_markup"`
}
