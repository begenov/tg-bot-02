package handlers

import (
	"github.com/begenov/tg-bot-02/internal/api"
	"github.com/begenov/tg-bot-02/internal/services"
)

type Handlers struct {
	services    *services.Service
	telegramAPI *api.TelegramAPI
}

func NewHadnler(service *services.Service, api *api.TelegramAPI) *Handlers {
	return &Handlers{
		services:    service,
		telegramAPI: api,
	}
}

func (h *Handlers) Start() error {
	updates := h.telegramAPI.GetUpdates()
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.Message.IsCommand() {
			if err := h.commandHandle(update.Message); err != nil {
				return err
			}
			continue
		}
		if err := h.handleMessage(update.Message); err != nil {
			continue
		}
	}
	return nil
}
