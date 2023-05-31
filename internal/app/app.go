package app

import (
	"github.com/joho/godotenv"
	"log"
	"main/internal/config"
	"main/internal/errors"
	"main/internal/speech_bot"
	"main/pkg/bot_middlewares"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func RunApp() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Recover")
			log.Println(r)
		}
	}()

	if _, err := os.Stat("voices"); os.IsNotExist(err) {
		log.Println("Created directory voices.")

		err = os.Mkdir("voices", 0777)
		if err != nil {
			log.Println(err)
		}
	}

	conf := config.New()

	var err error = nil
	if conf.TelegramBot.Token == "" {
		err = errors.NewSpeechBotError("Not found TELEGRAM_BOT_TOKEN in environment")
		log.Println(err)
	}
	if conf.YandexApi.Key == "" {
		err = errors.NewSpeechBotError("Not found YANDEX_API_KEY in environment")
		log.Println(err)
	}
	if err != nil {
		log.Println("Exit")
		return
	}

	speechBot := speech_bot.NewSpeechBot(conf.TelegramBot.Token, conf.YandexApi.Key, true)
	speechBot.AddMiddleware(bot_middlewares.GlebKhalitovBan)
	speechBot.Run()
}
