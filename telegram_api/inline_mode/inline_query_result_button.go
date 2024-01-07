package inline

import "github.com/sudak-91/telegrambotgo/telegram_api/types"

type InlineQueryResultButton struct {
	Text          string            `json:"text"`
	WebApp        *types.WebAppInfo `json:"web_app,omitempty"`
	StartParametr string            `json:"start_parametr,omitempty"`
}
