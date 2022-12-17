package inline

import types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"

type InlineQueryResultAudio struct {
	Type                string                              `json:"type"`
	ID                  string                              `json:"id"`
	AudioUrl            string                              `json:"audio_url"`
	Title               string                              `json:"title"`
	Caption             string                              `json:"caption,omitempty"`
	ParseMode           string                              `json:"parse_mode,omitempty"`
	CaptionEntities     []*types.TelegramMessageEntity      `json:"caption_entities,omitempty"`
	Performer           string                              `json:"performer,omitepty"`
	AudioDuration       int32                               `json:"audio_duration,omitempty"`
	ReplyMarkup         *types.TelegramInlineKeyboardMarkup `json:"reply_markup"`
	InputMessageContent *InputMessageContent                `json:"input_message_content"`
}