package inline

type InputLocationMessageContent struct {
	Latitude             float32 `json:"latitude"`
	Longitude            float32 `json:"longitude"`
	HorizontalAccuracy   float32 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int32   `json:"live_period,omitempty"`
	Heading              int32   `json:"heading,omitempty"`
	ProximityAlertRadius int32   `json:"proximity_alert_radius,omitempty"`
}
