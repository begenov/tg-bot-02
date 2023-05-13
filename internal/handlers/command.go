package handlers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (h *Handlers) commandHandle(message *tgbotapi.Message) error {
	switch message.Command() {
	case "/start":
		return h.startHandleCommand(message)
	default:
		return h.unknownHandleCommand(message)
	}
}
