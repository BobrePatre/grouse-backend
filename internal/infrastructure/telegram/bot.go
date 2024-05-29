package telegram

import (
	"github.com/PaulSonOfLars/gotgbot/v2"
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"log/slog"
)

type Config struct {
	BotToken string `json:"botToken" env:"BOT_TOKEN" validate:"required""`
}

func LoadConfig(validator *validator.Validate) (*Config, error) {
	var cfg struct {
		Config Config `json:"telegram" env-prefix:"TELEGRAM_"`
	}
	err := cleanenv.ReadConfig("config.json", &cfg)
	if err != nil {
		err := cleanenv.ReadEnv(&cfg)
		if err != nil {
			return nil, err
		}
	}
	err = validator.Struct(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg.Config, nil
}

func NewBot(cfg *Config, logger *slog.Logger) (*gotgbot.Bot, error) {
	return gotgbot.NewBot(cfg.BotToken, nil)
}
