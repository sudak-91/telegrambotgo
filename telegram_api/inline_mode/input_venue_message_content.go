package inline

type InputVenueMessageContent struct {
	Latitude        float32 `json:"latitude"`
	Longitude       float32 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareId    string  `json:"foursquare_id,omitempty"`
	GooglePlaceId   string  `json:"google_place_id,omitempt"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
}
