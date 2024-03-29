package inline

import types "github.com/sudak-91/telegrambotgo/telegram_api/types"

type InlineQueryResultLocation struct {
	Type                 string                      `json:"type"`
	ID                   string                      `json:"id"`
	Latitude             float32                     `json:"latitude"`
	Longitude            float32                     `json:"longitude"`
	Title                string                      `json:"title"`
	HorizontalAccuracy   float32                     `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int32                       `json:"live_period,omitempty"`
	Heading              int32                       `json:"heading,omitempty"`
	ProximityAlertRadius int32                       `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent  *InputMessageContent        `json:"in[ut_message_content"`
	ThumbUrl             string                      `json:"thimb_url,omitempty"`
	ThumbWidth           int32                       `json:"thumb_width,omitempty"`
	ThumbHeight          int32                       `json:"thumb_height,omitempty"`
}
