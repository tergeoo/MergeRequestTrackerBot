package main

import (
	"MRTrackerBot/model"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"net/http"
)

const botToken = "7116769840:AAH5xYFDYiXGq5j8PmBq-_cckvwfTMLR37o"
const chatID = "-1002137926837"

var devs = make(map[string]string)

var bot, botError = tgbotapi.NewBotAPI(botToken)

func formatMessage(webhook model.Webhook) string {
	var author string

	if shouldTagAuthor(webhook) {
		author = fmt.Sprintf("Author: @%s\n", devs[webhook.User.Username])
	} else {
		author = fmt.Sprintf("Author: %s\n", devs[webhook.User.Username])
	}

	link := fmt.Sprintf("[%s](%s)\n", webhook.ObjectAttributes.Title, webhook.ObjectAttributes.URL)
	status := fmt.Sprintf("Status: %s\n", webhook.ObjectAttributes.State)
	message := fmt.Sprintf("%s%s%s", link, status, author)

	return message
}

func shouldTagAuthor(webhook model.Webhook) bool {
	isApproved := webhook.ObjectAttributes.DetailedMergeStatus == "approved"
	hasComments := len(webhook.ObjectAttributes.Note) > 0

	return isApproved || hasComments
}

func sendMessage(w http.ResponseWriter, webhook model.Webhook) {
	if botError != nil {
		log.Panic(botError)
	}

	message := tgbotapi.NewMessageToChannel(chatID, formatMessage(webhook))
	message.ParseMode = "Markdown"
	result, err := bot.Send(message)
	if err != nil {
		log.Println("Error sending message:", err)
		http.Error(w, "Failed to send message", http.StatusInternalServerError)
	}

	fmt.Println(result.MessageID) //TODO потом сохранять в файл чтобы изменять

	log.Println("message sent")

	mrMessage := model.MergeRequestMessage{
		MessageId:      result.MessageID,
		MergeRequestId: webhook.ObjectAttributes.ID,
		Title:          webhook.ObjectAttributes.Title,
		Link:           webhook.ObjectAttributes.URL,
		Status:         webhook.ObjectAttributes.State,
		AuthorUserName: webhook.User.Username,
		HasComments:    len(webhook.ObjectAttributes.Note) > 0,
	}

	saveMessage(mrMessage)
}

func findMessage() {
	//message := tgbotapi.UpdatesChannel()
}

func initDevelopers() {
	devs = map[string]string{
		"g.ter-grigoryants": "tergeo",
		"r.sharafutdinov":   "BullyBoo",
		"n.korovkin":        "nkorovkin",
		"m.makarov":         "deufelil",
	}
}
