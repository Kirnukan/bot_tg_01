package commands

import (
	"github.com/Kirnukan/bot_tg_01/internal/service/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Commander struct {
	bot           *tgbotapi.BotAPI
	entityService *entity.Service
}

func NewCommander(
	bot *tgbotapi.BotAPI,
	entityService *entity.Service,
) *Commander {
	return &Commander{
		bot:           bot,
		entityService: entityService,
	}
}
