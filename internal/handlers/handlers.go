package handlers

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handlers) handleMessage(message *tgbotapi.Message) error {
	return nil
}

func (h *Handlers) startHandleCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Здравствуйте, это Telegram-bot по поиску работы и сотрудников.\nВыберите язык:")
	inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("казахский", "kazakh"),
			tgbotapi.NewInlineKeyboardButtonData("русский", "russian"),
		),
	)
	msg.ReplyMarkup = inlineKeyboard

	h.telegramAPI.Send(msg)
	return nil
}

func (h *Handlers) unknownHandleCommand(message *tgbotapi.Message) error {
	return nil
}
