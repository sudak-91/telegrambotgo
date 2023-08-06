package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type SendMessageResult struct {
	Ok     bool          `json:"ok"`
	Result types.Message `json:"result"`
}

type SendMessage struct {
	ChatID                   int64                       `json:"chat_id"`
	Text                     string                      `json:"text"`
	ParseMode                string                      `json:"parse_mode,omitempty"`
	Entities                 []types.MessageEntity       `json:"entities,omitempty"`
	DisableWebPagePreview    bool                        `json:"disable_web_page_preview,omitempty"`
	DisableNotification      bool                        `json:"disable_notification,omitempty"`
	ProtectContent           bool                        `json:"protect_content,omitempty"`
	ReplyToMessageID         int32                       `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply bool                        `json:"allow_sending_without_reply,omitempty"`
	ReplayMarkup             *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func SendMessageMethod(SecretKey string, Message SendMessage) error {
	URL := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", SecretKey)
	postBody, err := json.Marshal(Message)
	if err != nil {
		return err
	}
	Body := bytes.NewBuffer(postBody)
	resp, err := http.Post(URL, "application/json", Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	var Result SendMessageResult
	if err = json.Unmarshal(respBody, &Result); err != nil {
		return err
	}
	if !Result.Ok {
		return fmt.Errorf("Has erros: %s", respBody)
	}
	return nil

}
