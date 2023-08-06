package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	types "github.com/sudak-91/telegrambotgo/telegram_api/types"
)

type SendDocumentResult struct {
	Ok     bool          `json:"ok"`
	Result types.Message `json:"result"`
}

type SendDocument struct {
	ChatId                      int64                       `json:"chat_id"`
	Document                    string                      `json:"document"`
	Thumb                       string                      `json:"thumb,omitempty"`
	Caption                     string                      `json:"caption,omitempty"`
	ParseMode                   string                      `json:"parse_mode,omitempty"`
	CaptionEntity               *[]types.MessageEntity      `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool                        `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool                        `json:"disable_notification,omitepmpty"`
	ProtectContent              bool                        `json:"protect_content,omitempty"`
	ReplyToMessageID            int32                       `json:"reply_to_message_id,omitempty"`
	AllowSendingWithoutReply    bool                        `json:"allow_sending_without_reply,omitempty"`
	ReplyMarkup                 *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func SendDocumentMethod(SecretKey string, Document SendDocument) error {
	URL := fmt.Sprintf("https://api.telegram.org/bot%s/sendDocument", SecretKey)
	postBody, err := json.Marshal(Document)
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
	var Result SendDocumentResult
	if err = json.Unmarshal(respBody, &Result); err != nil {
		return err
	}
	log.Printf("%v", Result)
	if !Result.Ok {
		return fmt.Errorf("Has erros: %s", respBody)
	}
	return nil
}
