package types

type TelegramReplayMarkupKeyboard struct {
	Keyboard              [][]TelegramKeyboardButton `json:"keyboard"`
	ResizeKeyboard        bool                       `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       bool                       `json:"one_time_keuboard,omitempty"`
	InputFieldPlaceholder string                     `json:"input_field_placeholder,omitempty"`
	Selective             bool                       `json:"selective,omitempty"`
}
