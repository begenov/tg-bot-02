package app

import (
	"github.com/begenov/tg-bot-02/config"
	"github.com/begenov/tg-bot-02/db"
	"github.com/begenov/tg-bot-02/internal/api"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Run(cfg *config.Config) error {
	bot, err := tgbotapi.NewBotAPI(cfg.TelegramAPI.Token)

	if err != nil {
		return err
	}

	bot.Debug = true

	telegrmaAPI := api.NewTelegramAPI(bot)

	db, err := db.NewDB(cfg.DB.Driver, cfg.DB.DSN)

	if err != nil {
		return err
	}

	return nil
}
