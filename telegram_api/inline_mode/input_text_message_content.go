package inline

import types "github.com/sudak-91/telegrambotgo/telegram_api/types"

type InputTextMessageContent struct {
	MessageText           string                 `json:"message_text"`
	ParseMode             string                 `json:"parse_mode,omitepmty"`
	Entities              []*types.MessageEntity `json:"entities,omitempty"`
	DisableWebPagePreview bool                   `json:"disable_web_page_preview,omitempty"`
}
