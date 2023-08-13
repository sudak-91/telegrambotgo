package types

type MenuButtonWebApp struct {
	Text   string      `json:"text"`
	WebApp *WebAppInfo `json:"web_app"`
}
