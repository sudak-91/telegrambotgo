package types

type TelegramVenue struct {
	Location        TelegramLocation `json:"location"`
	Title           string           `json:"title"`
	Address         string           `json:"address"`
	FoursquareID    string           `json:"foursquare_id,omitempty"`
	Foursquaretype  string           `json:"foursquare_type,omitempty"`
	GooglePlaceID   string           `json:"google_place_id,omitempty"`
	GooglePlaceType string           `json:"google_place_type,omitempty"`
}
