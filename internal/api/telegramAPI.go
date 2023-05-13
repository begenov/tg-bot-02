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

func (api *TelegramAPI) GetUpdates() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)

	u.Timeout = 60

	updates := api.bot.GetUpdatesChan(u)

	return updates
}

func (api *TelegramAPI) Send(c tgbotapi.Chattable) (tgbotapi.Message, error) {
	return api.bot.Send(c)
}
