package types

type TelegramUserProfilePhotos struct {
	TotalCount int32          `json:"total_count"`
	Photos     [][]*PhotoSize `json:"photos"`
}
