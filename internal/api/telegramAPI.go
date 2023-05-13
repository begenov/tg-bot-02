package api

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type TelegramAPI struct {
	bot *tgbotapi.BotAPI
}

func NewTelegramAPI(bot *tgbotapi.BotAPI) *TelegramAPI {
	return &TelegramAPI{
		bot: bot,
	}
}
