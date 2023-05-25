package main

import (
	"fmt"
	"github.com/Kirnukan/bot_tg_01/internal/service/entity"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load("local.env")
	dotEnvApi := os.Getenv("API_KEY")
	bot, err := tgbotapi.NewBotAPI(dotEnvApi)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.UpdateConfig{
		Timeout: 60,
	}

	updates := bot.GetUpdatesChan(u)

	entityService := entity.NewService()

	commander := NewCommander(bot, entityService)

	for update := range updates {
		if update.Message != nil {

			switch update.Message.Command() {
			case "help":
				commander.Help(update.Message)
			case "hello":
				commander.Hello(update.Message)
			case "list":
				commander.List(update.Message)
			default:
				commander.Default(update.Message)
			}
		}
	}
}

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

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list entities",
	)

	c.bot.Send(msg)
}

func (c *Commander) Hello(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Hello, "+inputMessage.From.UserName)

	c.bot.Send(msg)
}

func (c *Commander) List(inputMessage *tgbotapi.Message) {
	entitiesMsgText := "Entities List: \n\n"

	entities := c.entityService.List()

	for i, e := range entities {
		entitiesMsgText += fmt.Sprintf("%d) %s\n", i+1, e.Title)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, entitiesMsgText)

	c.bot.Send(msg)
}

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	//msg.ReplyToMessageID = update.Message.MessageID

	c.bot.Send(msg)
}
