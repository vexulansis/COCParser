package tg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vexulansis/COCParser/pkg/api"
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
			b.Respond(&update)
			if err != nil {
				return err
			}
		}
	}
}
func (b *TGBot) Respond(u *Update) error {
	msg := new(BotMessage)
	msg.Chat_id = u.Message.Chat.Chat_id
	clan, err := api.GetClanByTag(u.Message.Text)
	if err != nil {
		log.Fatal()
	}
	msg.Text = clan.Name
	_, err = b.sendMessage(msg)
	if err != nil {
		return err
	}
	return nil
}
