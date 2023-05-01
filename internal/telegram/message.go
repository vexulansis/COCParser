package tg

// https://core.telegram.org/bots/api#message
type Message struct {
	Text       string `json:"text"`
	Message_id int    `json:"message_id"`
	Chat       Chat   `json:"chat"`
}

type BotMessage struct {
	Text       string `json:"text"`
	Message_id int    `json:"message_id"`
	Chat_id    int    `json:"chat_id"`
}
