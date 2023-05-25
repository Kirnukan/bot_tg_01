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

	for update := range updates {
		if update.Message != nil {

			switch update.Message.Command() {
			case "help":
				helpCommand(bot, update.Message)
			case "hello":
				helloCommand(bot, update.Message)
			case "list":
				listCommand(bot, update.Message, entityService)
			default:
				defaultMessaging(bot, update.Message)
			}
		}
	}
}

func helpCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list entities",
	)

	bot.Send(msg)
}

func helloCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Hello, "+inputMessage.From.UserName)

	bot.Send(msg)
}

func listCommand(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, entityService *entity.Service) {
	entitiesMsgText := "Entities List: \n\n"

	entities := entityService.List()

	for i, e := range entities {
		entitiesMsgText += fmt.Sprintf("%d) %s\n", i+1, e.Title)
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, entitiesMsgText)

	bot.Send(msg)
}

func defaultMessaging(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	log.Printf("[%s] %s", inputMessage.From.UserName, inputMessage.Text)
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "You wrote: "+inputMessage.Text)

	//msg.ReplyToMessageID = update.Message.MessageID

	bot.Send(msg)
}
