package app

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"main/internal/bot"
	"main/internal/config"
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
			fmt.Println(r)
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
	bot.RunSpeechBot(conf.TelegramBot.Token, conf.YandexApi.Token, true)
}
