package main

import (
	"fmt"
	tgbotapi "github.com/crocone/tg-bot"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var bot *tgbotapi.BotAPI

var startMenu = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Say hello", "hi"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Say bye", "bye"),
	),
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("env not load")
	}
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_API_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Something with API_KEY: %v", err)
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	if err != nil {
		log.Fatalf("failed to listening message: %v", err)
	}
	for update := range updates {
		log.Println(update)
		if update.CallbackQuery != nil {
			callbacks(update, bot)
		} else if update.Message.IsCommand() {
			commands(update, bot)
		} else {
			log.Println("au not working")
		}
	}
}

func callbacks(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	data := update.CallbackQuery.Data
	chatID := update.CallbackQuery.From.ID
	firstName := update.CallbackQuery.From.FirstName
	lastName := update.CallbackQuery.From.LastName
	var text string
	switch data {
	case "hi":
		text = fmt.Sprintf("Hello %v %v", firstName, lastName)
	case "bye":
		text = fmt.Sprintf("Bye %v %v", firstName, lastName)
	default:
		text = "invalid message or null"

	}
	msg := tgbotapi.NewMessage(chatID, text)
	bot.Send(msg)

}

func commands(update tgbotapi.Update, bot *tgbotapi.BotAPI) {
	command := update.Message.Command()
	switch command {
	case "start":
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Choice actions")
		msg.ReplyMarkup = startMenu
		msg.ParseMode = "Markdown"
		ok, err := bot.Send(msg)
		if err != nil {
			log.Printf("Не удалось отправить сообщение: %v", err)
		}
		log.Println(ok)
	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "invalid command")
		bot.Send(msg)
	}

}

//func sendMessage(msg tgbotapi.Chattable) {
//	ok, err := bot.Send(msg)
//	if err != nil {
//		log.Fatalf("Произошла ошибка при отправке сообщения: %v", err)
//	}
//	log.Println(ok)
//}
