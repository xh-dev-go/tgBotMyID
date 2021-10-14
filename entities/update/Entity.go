package update

type Response struct {
	Ok bool `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	UpdateId int64 `json:"update_id"`
	Message Message `json:"message"`
}

type Message struct {
	MessageId int64 `json:"message_id"`
	From From `json:"from"`
}

type From struct {
	Id int64 `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	LanguageCode string `json:"language_code"`
}
