package app

import (
	"github.com/begenov/tg-bot-02/config"
	"github.com/begenov/tg-bot-02/db"
	"github.com/begenov/tg-bot-02/internal/api"
	"github.com/begenov/tg-bot-02/internal/handlers"
	"github.com/begenov/tg-bot-02/internal/repository"
	"github.com/begenov/tg-bot-02/internal/services"
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

	repository := repository.NewRepository(db)

	services := services.NewService(repository)

	handler := handlers.NewHadnler(services, telegrmaAPI)

	err = handler.Start()
	if err != nil {
		return err
	}
	return nil
}
