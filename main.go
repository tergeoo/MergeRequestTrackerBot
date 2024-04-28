//package main

//
//func main() {
//	if err := run(context.Background()); err != nil {
//		log.Fatal(err)
//	}

//port := "8080"
//log.Printf("Starting server on port %s", port)
//log.Fatal(http.ListenAndServe(":"+port, nil))
//}

//func run(ctx context.Context) error {
//	fmt.Println("run start")
//	listener, err := ngrok.Listen(ctx,
//		config.HTTPEndpoint(),
//		ngrok.WithAuthtokenFromEnv(),
//	)
//
//	if err != nil {
//		fmt.Println("run error")
//		return err
//	}
//
//	fmt.Println(listener.URL())
//	log.Println("App URL", listener.URL())
//	return http.Serve(listener, http.HandlerFunc(handler))
//}
//
//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(w, "<h1>Hello from ngrok-go!</h1>")
//}
//

package main

import (
	"MRTrackerBot/model"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.ngrok.com/ngrok"
	"golang.ngrok.com/ngrok/config"
)

func main() {
	go initDevelopers()
	go createFileIfNotExist()

	if err := run(context.Background()); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	ln, err := ngrok.Listen(ctx,
		config.HTTPEndpoint(),
		ngrok.WithAuthtokenFromEnv(),
	)

	if err != nil {
		return err
	}

	log.Println("Ingress established at:", ln.URL())

	http.HandleFunc("/webhook", handleWebhook)

	return http.Serve(ln, nil)
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
	err := json.Unmarshal(body, &webhook)

	if err != nil {
		log.Println("Error decoding webhook payload:", err)
		http.Error(w, "Invalid payload", http.StatusBadRequest)
		return
	}

	sendMessage(w, webhook)
}
