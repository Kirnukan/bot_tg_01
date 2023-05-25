package commands

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	entitiesMsgText := "Entities List: \n\n"

	entities := c.entityService.List()

	for i, e := range entities {
		entitiesMsgText += fmt.Sprintf("%d) %s\n", i+1, e.Title)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, entitiesMsgText)

	c.bot.Send(msg)
}

func init() {
	registeredCommands["list"] = (*Commander).List
}
