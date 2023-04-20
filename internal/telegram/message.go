package tg

// https://core.telegram.org/bots/api#message
type Message struct {
	Message_id int    `json:"message_id"`
	Text       string `json:"text"`
	Chat       Chat   `json:"chat"`
}

type BotMessage struct {
	Message_id int    `json:"message_id"`
	Chat_id    int    `json:"chat_id"`
	Text       string `json:"text"`
}
