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

type GetChatMember struct {
	ChatID string `json:"chat_id"` //in format @channelname
	UserID int64  `json:"user_id"`
}

type GetChatMemberResult struct {
	Ok          bool              `json:"OK"`
	Result      *types.ChatMember `json:"result,omitempty"`
	Status      string            `json:"status"`
	Description string            `json:"description,omitempty"`
}

func GetChanMemberMethod(ChatID string, UserID int64, Key string) (types.ChatMember, error) {
	var ChatMember GetChatMember
	ChatMember.ChatID = ChatID
	ChatMember.UserID = UserID

	URL := fmt.Sprintf("https://api.telegram.org/bot%s/getChatMember", Key)
	postBody, err := json.Marshal(ChatMember)
	if err != nil {
		return types.ChatMember{}, err
	}
	Body := bytes.NewBuffer(postBody)

	resp, err := http.Post(URL, "application/json", Body)
	if err != nil {
		return types.ChatMember{}, err
	}
	defer resp.Body.Close()
	var Result GetChatMemberResult
	respbody, err := ioutil.ReadAll(resp.Body)
	log.Printf("GetChanMember: %s", respbody)
	if err != nil {
		return types.ChatMember{}, err
	}
	err = json.Unmarshal(respbody, &Result)
	log.Println(Result)
	if err != nil {
		return types.ChatMember{}, err
	}
	if Result.Ok != true {

		return types.ChatMember{}, fmt.Errorf("Has error: %s", Result.Description)

	}
	return *Result.Result, nil

}
