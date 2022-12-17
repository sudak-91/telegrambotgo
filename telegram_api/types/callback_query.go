package types

/*
https://core.telegram.org/bots/api#callbackquery
*/

type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"` //Sender
	Message         *Message `json:"message,omitempty"`
	InlineMessageId string   `json:"inline_message_id,omitempty"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data,omitempty"`
	GameShortName   string   `json:"game_short_name,omitempty"`
}
