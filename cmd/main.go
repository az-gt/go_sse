package main

import (
	"fmt"
	"log"
	"main/internal/events"
	"main/internal/webhooks"
	"net/http"
)

func main() {
	http.HandleFunc("/events", events.SendEvents)
	http.HandleFunc("/events/{ID}", events.SendEventsID)
	http.HandleFunc("/webhook", webhooks.HandleWebhook)
	http.HandleFunc("/webhook/{ID}", webhooks.HandleWebhookID)
	fmt.Println("Server running on port http://localhost:8060")
	log.Fatal(http.ListenAndServe(":8060", nil))
}
