package types

type ProximityAlertTrigered struct {
	Traveler *User `json:"traveler"`
	Watcher  *User `json:"watcher"`
	Distance int32 `json:"distance"`
}
