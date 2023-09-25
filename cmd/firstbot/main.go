package main

import (
	"firstbot/internal/service/product"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	godotenv.Load()
	token := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	productService := product.NewProductService()

	bot.Debug = true
	log.Printf("Authorized as account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			log.Printf("User %v, SEND: %s", update.Message.Chat.ID, update.Message.Text)

			switch update.Message.Command() {

			case "help":
				getHelp(bot, update.Message)
			case "list":
				getProductList(bot, update.Message, productService)
			}
		}
	}
}

func getHelp(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - product list",
	)
	bot.Send(msg)
}

func getProductList(bot *tgbotapi.BotAPI, inputMessage *tgbotapi.Message, service *product.ProductService) {
	for _, productL := range service.List() {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID,
			"Your product: "+productL.Title,
		)
		bot.Send(msg)

	}

}
