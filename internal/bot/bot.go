package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"main/pkg/download_file"
	"main/pkg/speech_to_text"
)

func RunSpeechBot(token string, apiToken string, noneStop bool) {
	var err error

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Println(err)
	}

	if noneStop {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
				bot.StopReceivingUpdates()
				RunSpeechBot(token, apiToken, true)
			}
		}()
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Voice != nil {

			file, _ := bot.GetFile(tgbotapi.FileConfig{
				FileID: update.Message.Voice.FileID,
			})

			filePath := "voices/" + update.Message.Voice.FileUniqueID + ".ogg"
			download_file.DownloadFile(file.Link(token), filePath)

			answer := speech_to_text.SpeechToText(apiToken, filePath)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
			msg.ReplyToMessageID = update.Message.MessageID
			bot.Send(msg)
		}
	}
}
