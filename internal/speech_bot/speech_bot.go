package speech_bot

import (
	"encoding/json"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"main/pkg/download_file"
	"main/pkg/speech_to_text"
)

func NewSpeechBot(token string, apiToken string, noneStop bool) *SpeechBot {
	return &SpeechBot{
		token:    token,
		apiToken: apiToken,
		noneStop: noneStop,
	}
}

type SpeechBot struct {
	token       string
	apiToken    string
	noneStop    bool
	middlewares []func(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error
}

func (sb *SpeechBot) Run() {
	var err error

	bot, err := tgbotapi.NewBotAPI(sb.token)
	if err != nil {
		log.Println(err)
	}

	if sb.noneStop {
		defer func() {
			if r := recover(); r != nil {
				log.Println(r)
				bot.StopReceivingUpdates()
				sb.Run()
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

		err := sb.runMiddlewares(bot, update.Message)
		if err != nil {
			continue
		}

		if update.Message.Voice != nil {
			file, _ := bot.GetFile(tgbotapi.FileConfig{
				FileID: update.Message.Voice.FileID,
			})

			filePath := "voices/" + update.Message.Voice.FileUniqueID + ".ogg"
			download_file.DownloadFile(file.Link(sb.token), filePath)

			answer := speech_to_text.SpeechToText(sb.apiToken, filePath)

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, answer)
			msg.ReplyToMessageID = update.Message.MessageID

			logSpeechMessage(update.Message, answer)
			bot.Send(msg)
		}
	}
}

func (sb *SpeechBot) runMiddlewares(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error {
	for _, v := range sb.middlewares {
		err := v(bot, message)
		if err != nil {
			return err
		}
	}
	return nil
}

func logSpeechMessage(message *tgbotapi.Message, VoiceText string) {
	jsonMessage, err := json.Marshal(message)
	if err == nil {
		log.Println(string(jsonMessage))
		log.Println(VoiceText)
	}
}

func (sb *SpeechBot) AddMiddleware(middleware func(bot *tgbotapi.BotAPI, message *tgbotapi.Message) error) {
	sb.middlewares = append(sb.middlewares, middleware)
}
