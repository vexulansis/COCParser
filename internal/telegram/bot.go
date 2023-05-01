package tg

import (
	"os"

	"github.com/joho/godotenv"
)

type TGBot struct {
	Token string
}

func NewBot() (*TGBot, error) {
	bot := new(TGBot)
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}
	bot.Token = os.Getenv("TGBOT_TOKEN")
	return bot, nil
}

func (b *TGBot) Start() error {
	offset := 0
	for {
		updates, err := b.getUpdates(offset)
		if err != nil {
			return err
		}
		for _, update := range updates {
			offset = update.Update_id + 1
			if err := b.Respond(&update); err != nil {
				return err
			}
		}
	}
}
func (b *TGBot) Respond(u *Update) error {
	msg := new(BotMessage)
	msg.Chat_id = u.Message.Chat.Chat_id

	return nil
}
