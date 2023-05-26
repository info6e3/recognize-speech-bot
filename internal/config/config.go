package config

import (
	"os"
)

type SqlConfig struct {
	Driver   string
	Host     string
	User     string
	Password string
	DbName   string
	SslMode  string
}

type TelegramBotConfig struct {
	Token string
}

type YandexApiConfig struct {
	Token string
}

type Config struct {
	Sql         SqlConfig
	TelegramBot TelegramBotConfig
	YandexApi   YandexApiConfig
}

// New returns a new Config struct
func New() Config {
	return Config{
		Sql: SqlConfig{
			Driver:   getEnv("SQL_DRIVER", ""),
			Host:     getEnv("SQL_HOST", ""),
			User:     getEnv("SQL_USER", ""),
			Password: getEnv("SQL_PASSWORD", ""),
			DbName:   getEnv("SQL_DB_NAME", ""),
			SslMode:  getEnv("SQL_SSL_MODE", ""),
		},
		TelegramBot: TelegramBotConfig{
			Token: getEnv("TELEGRAM_BOT_TOKEN", ""),
		},
		YandexApi: YandexApiConfig{
			Token: getEnv("YANDEX_API_KEY", ""),
		},
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}
