package utils

import (
	"MRTrackerBot/model"
	json2 "encoding/json"
	"fmt"
	"log"
	"os"
)

const messagesFileName = "files/message.json"
const developersFileName = "files/config.json"

func SaveMessage(mergeRequestId, messageId int) {
	saved := ReadSavedMessages()
	saved.MessageIds[mergeRequestId] = messageId

	saveJSON(messagesFileName, saved)
}

func ReadSavedMessages() model.Saved {
	file := readFile(messagesFileName)
	var saved model.Saved

	err := json2.Unmarshal(file, &saved)
	if err != nil {
		panic(err)
	}

	return saved
}

func ReadConfig() model.Config {
	file := readFile(developersFileName)
	var devs model.Config

	err := json2.Unmarshal(file, &devs)
	if err != nil {
		panic(err)
	}

	return devs
}

func CreateMessageFileIfNotExist() {
	_, err := os.ReadFile(messagesFileName)
	if err == nil {
		return
	}

	_, err = os.Create(messagesFileName)
	if err != nil {
		panic(err)
	}

	saveJSON(messagesFileName, model.Saved{MessageIds: map[int]int{}})
}

func CreateDevsFileIfNotExist() {
	_, err := os.ReadFile(developersFileName)
	if err == nil {
		return
	}

	_, err = os.Create(developersFileName)
	if err != nil {
		panic(err)
	}

	saveJSON(developersFileName, make(map[string]string))
}

func readFile(fileName string) []byte {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return file
}

func saveJSON(filename string, data interface{}) {
	fileData, err := json2.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
	}
	err = os.WriteFile(filename, fileData, 0644)
	if err != nil {
		log.Fatalf("Error writing file: %s", err.Error())
	}
	fmt.Println("Data successfully saved to", filename)
}
