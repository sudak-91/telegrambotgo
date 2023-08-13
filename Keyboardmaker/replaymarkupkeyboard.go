package keyboardmaker

import (
	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type ReplayMarkupKeyboard struct {
	KeyboardMarkup types.ReplayMarkupKeyboard
}

func (k *ReplayMarkupKeyboard) MakeGrid(raw int, column int) {
	k.KeyboardMarkup.Keyboard = make([][]types.KeyboardButton, raw)
	for i := range k.KeyboardMarkup.Keyboard {
		k.KeyboardMarkup.Keyboard[i] = make([]types.KeyboardButton, column)
	}

}
func (k *ReplayMarkupKeyboard) AddButton(text string, raw int, column int) {
	var Button types.KeyboardButton
	Button.Text = text
	k.KeyboardMarkup.Keyboard[raw][column] = Button
}
func (k ReplayMarkupKeyboard) ClearButton(raw int, column int) {
	return
}
func (k *ReplayMarkupKeyboard) GetKeyboard() types.ReplayMarkupKeyboard {
	return k.KeyboardMarkup
}
