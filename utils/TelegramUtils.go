package utils

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"strconv"
)

var botToken = ""
var chatID = -1
var bot *tgbotapi.BotAPI = nil

func SendMessage(text string) tgbotapi.Message {
	if len(botToken) == 0 || chatID == -1 {
		initBot()
	}

	message := tgbotapi.NewMessageToChannel(strconv.Itoa(chatID), text)
	message.ParseMode = "Markdown"
	result, err := bot.Send(message)
	if err != nil {
		log.Println("Error sending message:", err)
		panic(err)
	}

	return result
}

func UpdateTelegramMessage(messageId int, text string) {
	message := tgbotapi.NewEditMessageText(int64(chatID), messageId, text)
	message.ParseMode = "Markdown"

	_, err := bot.Send(message)
	if err != nil {
		log.Println("Error sending message:", err)
		panic(err)
	}
}

func initBot() {
	config := ReadConfig()
	chatID = config.ChaId
	botToken = config.BotToken

	var err error
	bot, err = tgbotapi.NewBotAPI(botToken) // Теперь здесь используется глобальная переменная
	if err != nil {
		panic(err)
	}
}
