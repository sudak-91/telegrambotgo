package methods

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GetChatMemberCount struct {
	ChatID string `json:"chat_id"`
}

type GetChatMemberCountResult struct {
	Ok     bool `json:"ok"`
	Result int  `json:"result"`
}

func GetChatMemberCountMethod(ChatID string, SecreteKey string) (int, error) {
	var chatmembercount GetChatMemberCount
	chatmembercount.ChatID = ChatID
	path := fmt.Sprintf("https://api.telegram.org/bot%s/getChatMemberCount", SecreteKey)
	jsonBody, err := json.Marshal(chatmembercount)
	if err != nil {
		ErrMsg := fmt.Errorf("GetChatMember has error: %s", err.Error())
		return int(-1), ErrMsg
	}
	Body := bytes.NewBuffer(jsonBody)
	resp, err := http.Post(path, "application/json", Body)
	if err != nil {
		ErrMsg := fmt.Errorf("GetChatMember has error: %s", err.Error())
		return int(-1), ErrMsg
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ErrMsg := fmt.Errorf("GetChatMember has error: %s", err.Error())
		return int(-1), ErrMsg
	}
	var Result GetChatMemberCountResult
	json.Unmarshal(respBody, &Result)
	if err != nil {
		ErrMsg := fmt.Errorf("GetChatMember has error: %s", err.Error())
		return int(-1), ErrMsg
	}
	return Result.Result, nil

}
