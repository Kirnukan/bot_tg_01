package main

import (
	"github.com/Kirnukan/bot_tg_01/internal/app/commands"
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

	commander := commands.NewCommander(bot, entityService)

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
