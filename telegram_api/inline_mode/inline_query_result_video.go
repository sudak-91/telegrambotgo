package inline

import types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"

type InlineQueryResultVideo struct {
	Type                string                              `json:"type"`
	Id                  string                              `json:"id"`
	VideoURL            string                              `json:"video_url"`
	MimeType            string                              `json:"mime_type"`
	ThumbURL            string                              `json:"thumb_url"`
	Title               string                              `json:"title"`
	Caption             string                              `json:"caption,omitempty"`
	ParseMode           string                              `json:"parse_mode"`
	CaptionEntities     []*types.TelegramMessageEntity      `json:"caption_entities,omitempty"`
	VideoWidth          int32                               `json:"video_width,omitempty"`
	VideoHeight         int32                               `json:"video_height,omitempty"`
	VideoDuration       int32                               `json:"video_duration,omitempty"`
	Description         string                              `json:"descriptio,omitempty"`
	ReplyMarkup         *types.TelegramInlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent                `json:"input_message_content"`
}
