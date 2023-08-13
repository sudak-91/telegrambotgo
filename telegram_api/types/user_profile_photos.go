package types

type UserProfilePhotos struct {
	TotalCount int32          `json:"total_count"`
	Photos     [][]*PhotoSize `json:"photos"`
}
