package types

type MenuButton struct {
	Type string      `json:"type"`
	Data interface{} `json:"_,omitempt"`
}
