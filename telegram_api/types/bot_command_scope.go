package types

type BotCommandScope struct {
	Type string      `json:"type"`
	Data interface{} `json:"_,omitempty"`
}
