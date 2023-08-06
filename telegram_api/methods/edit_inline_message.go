package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"net/http"

	tgtypes "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type EditMessageText struct {
	ChatID                string                        `json:"chat_id,omitempty"`
	MessageID             int                           `json:"message_id,omitempty"`
	InlineMessageID       int                           `json:"inline_message_id,omitempty"`
	Text                  string                        `json:"text"`
	ParseMode             string                        `json:"parse_mode,omitempty"`
	Entities              []*tgtypes.MessageEntity      `json:"entites,omitempty"`
	DisableWebPagePreview bool                          `json:"disable_web_page_preview,omitempty"`
	ReplyMarkup           *tgtypes.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}
type EditMessageTextResult struct {
	Ok     bool `json:"ok"`
	Result *tgtypes.Message
}

func EditMessageTextMethod(EditMessage EditMessageText, Key string) error {
	URL := fmt.Sprintf("https://api.telegram.org/bot%s/editMessageText", Key)
	postBody, err := json.Marshal(EditMessage)
	if err != nil {
		return err
	}

	Body := bytes.NewBuffer(postBody)
	resp, err := http.Post(URL, "application/json", Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	var Result GetChatMemberResult
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(respBody, &Result); err != nil {
		return err
	}
	if !Result.Ok {
		return fmt.Errorf("Has error: %s", respBody)
	}
	return nil

}
