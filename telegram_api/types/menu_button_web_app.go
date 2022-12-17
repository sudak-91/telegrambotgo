package types

type MenuButtonWebApp struct {
	Text   string              `json:"text"`
	WebApp *TelegramWebAppInfo `json:"web_app"`
}
