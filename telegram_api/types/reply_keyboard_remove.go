package types

type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard,omitempty"`
	Selective      bool `json:"selective,omitempty"`
}
