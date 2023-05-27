package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
)

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	args := inputMessage.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("Wrong args", args)
		return
	}

	entity, err := c.entityService.Get(idx)
	if err != nil {
		log.Printf("Fail to get entity with idx %d: %v", idx, err)
		return
	}

	msg := tgbotapi.NewMessage(
		inputMessage.Chat.ID,
		entity.Title,
	)
	c.bot.Send(msg)
}

func init() {
	registeredCommands["get"] = (*Commander).Get
}
