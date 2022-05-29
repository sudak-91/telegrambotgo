package types

type TelegramProximityAlertTrigered struct {
	Traveler TelegramUser `json:"traveler"`
	Watcher  TelegramUser `json:"watcher"`
	Distance int32        `json:"distance"`
}
