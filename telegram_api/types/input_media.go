package types

type InputMedia struct {
	Type string      `json:"type"`
	Data interface{} `json"_,omitempty"`
}
