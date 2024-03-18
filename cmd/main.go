package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"main/internal/events"
	"main/internal/webhooks"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/events", events.SendEvents)
	http.HandleFunc("/events/{ID}", events.SendEventsID)
	http.HandleFunc("/webhook", webhooks.HandleWebhook)
	http.HandleFunc("/webhook/{ID}", webhooks.HandleWebhookID)

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	AppPort := os.Getenv("APP_PORT")

	fmt.Println("Server running on port http://localhost:" + AppPort)
	log.Fatal(http.ListenAndServe(":"+AppPort, nil))
}
