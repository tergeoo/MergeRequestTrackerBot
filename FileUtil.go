package main

import (
	"MRTrackerBot/model"
	json2 "encoding/json"
	"fmt"
	"log"
	"os"
)

const fileName = "message.json"

func saveMessage(message model.MergeRequestMessage) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	var saved model.Saved

	err = json2.Unmarshal(file, &saved)
	if err != nil {
		panic(err)
	}

	saved.Messages[message.MessageId] = message

	saveJSON(fileName, saved)
}

func createFileIfNotExist() {
	_, err := os.ReadFile(fileName)
	if err == nil {
		return
	}

	_, err = os.Create(fileName)
	if err != nil {
		panic(err)
	}

	saveJSON(fileName, model.Saved{Messages: map[int]model.MergeRequestMessage{}})
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
