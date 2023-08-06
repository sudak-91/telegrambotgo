package inline

import types "github.com/sudak-91/telegrambotgo/telegram_api/types"

type InlineQueryResultVenue struct {
	Type                string                      `json:"type"`
	ID                  string                      `json:"id"`
	Latitude            float32                     `json:"latitude"`
	Longitude           float32                     `json:"longitude"`
	Title               string                      `json:"title"`
	Address             string                      `json:"address"`
	FoursquareId        string                      `json:"foursquare_id,omitempty"`
	FoursquareType      string                      `json:"foursquare_type,omitempty"`
	GooglePlaceId       string                      `json:"google_place_id,omitempty"`
	GooglePlaceType     string                      `json:"google_place_type,omitempty"`
	ReplyMarkup         *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent *InputMessageContent        `json:"input_message_content,omitempty"`
	ThumbUrl            string                      `json:"thumb_url,omitempty"`
	ThumbWidth          int32                       `json:"thumb_width,omitempty"`
	ThumbHeight         int32                       `json:"thumb_height,omitepty"`
}
