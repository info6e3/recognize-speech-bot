package bot_middlewares

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
	"time"
)

const glebId = 692528838

func GlebKhalitovBan(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	if message.From.ID == glebId {
		if message.Voice != nil && poshelNah() {
			msg := tgbotapi.NewMessage(message.Chat.ID, "пошел ты нах*** Глеб Халитов")
			msg.ReplyToMessageID = message.MessageID
			bot.Send(msg)
		}
		return &GlebKhalitovError{}
	}
	return nil
}

func poshelNah() bool {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(10) > 0 {
		return false
	}
	return true
}

type GlebKhalitovError struct{}

func (*GlebKhalitovError) Error() string {
	return "Это Глеб Халитов!"
}
