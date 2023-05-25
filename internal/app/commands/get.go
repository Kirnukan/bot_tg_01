package commands

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (c *Commander) Get(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Your nickname is "+inputMessage.From.UserName)
	c.bot.Send(msg)
}

func init() {
	registeredCommands["get"] = (*Commander).Get
}
