package types

type TelegramUserProfilePhotos struct {
	TotalCount int32                  `json:"total_count"`
	Photos     [][]*TelegramPhotoSize `json:"photos"`
}
