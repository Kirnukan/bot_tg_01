package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Hello(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Hello, "+inputMessage.From.UserName)

	c.bot.Send(msg)
}

func init() {
	registeredCommands["hello"] = (*Commander).Hello
}
