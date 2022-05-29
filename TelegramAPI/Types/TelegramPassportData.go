package types

//@TODO
type TelegramPassportData struct {
	Data        []TelegramEncryptedPassportElement `json:"data"`
	Credentials TelegramEncryptedCredentials       `json:"credentials"`
}
