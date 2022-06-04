package keyboardmaker

import (
	types "github.com/sudak-91/telegrambotgo/TelegramAPI/Types"
)

type ReplayMarkupKeyboard struct {
	KeyboardMarkup types.TelegramReplayMarkupKeyboard
}

func (k *ReplayMarkupKeyboard) MakeGrid(raw int, column int) {
	k.KeyboardMarkup.Keyboard = make([][]types.TelegramKeyboardButton, raw)
	for i := range k.KeyboardMarkup.Keyboard {
		k.KeyboardMarkup.Keyboard[i] = make([]types.TelegramKeyboardButton, column)
	}

}
func (k *ReplayMarkupKeyboard) AddButton(text string, raw int, column int) {
	var Button types.TelegramKeyboardButton
	Button.Text = text
	k.KeyboardMarkup.Keyboard[raw][column] = Button
}
func (k ReplayMarkupKeyboard) ClearButton(raw int, column int) {
	return
}
func (k *ReplayMarkupKeyboard) GetKeyboard() types.TelegramReplayMarkupKeyboard {
	return k.KeyboardMarkup
}
