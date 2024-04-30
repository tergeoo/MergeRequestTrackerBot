package main

import (
	"MRTrackerBot/manager"
	"MRTrackerBot/model"
	"MRTrackerBot/utils"
	"context"
	json2 "encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Start")
	utils.CreateMessageFileIfNotExist()
	utils.CreateDevsFileIfNotExist()

	fmt.Println("run")
	if err := run(context.Background()); err != nil {
		fmt.Println("err")
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	//ln, err := ngrok.Listen(ctx,
	//	config.HTTPEndpoint(),
	//	ngrok.WithAuthtokenFromEnv(),
	//)
	//
	//if err != nil {
	//	return err
	//}
	//
	//log.Println("Ingress established at:", ln.URL())

	http.HandleFunc("/webhook", handleWebhook)

	return http.ListenAndServe("localhost:8080", nil)
	//return http.Serve(ln, nil)
}

func handleWebhook(w http.ResponseWriter, r *http.Request) {
	body, er := io.ReadAll(r.Body)
	if er != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}
	fmt.Println(r.URL)
	fmt.Println(string(body))

	var webhook model.Webhook
	err := json2.Unmarshal(body, &webhook)
	if err != nil {
		log.Println("Error decoding webhook payload:", err)
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	if webhook.ObjectKind == "merge_request" {
		manager.ProcessMergeRequest(webhook)
		return
	}

	if webhook.ObjectKind == "note" {
		manager.ProcessNote(webhook)
		return
	}
}
