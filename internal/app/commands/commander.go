package commands

import (
	"github.com/Kirnukan/bot_tg_01/internal/service/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

var registeredCommands = map[string]func(c *Commander, msg *tgbotapi.Message){}

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

func (c *Commander) HandleUpdate(update tgbotapi.Update) {
	defer func() {
		if panicValue := recover(); panicValue != nil {
			log.Printf("recovered from panic: %v", panicValue)
		}
	}()
	if update.Message != nil {

		command, ok := registeredCommands[update.Message.Command()]
		if ok {
			command(c, update.Message)
		} else {
			c.Default(update.Message)
		}
	}
}
