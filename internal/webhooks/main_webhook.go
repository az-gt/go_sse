package webhooks

import (
	"encoding/json"
	"log"
	"main/internal/events"
	"main/internal/models"
	"net/http"
)

func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	// Decodificar los datos del nuevo registro de la solicitud
	var newRecord models.Record
	err := json.NewDecoder(r.Body).Decode(&newRecord)
	if err != nil {
		log.Println("Error decoding webhook payload:", err)
		http.Error(w, "Error decoding webhook payload", http.StatusBadRequest)
		return
	}

	// Enviar los datos del nuevo registro a los clientes a través de Server-Sent Events
	sendToClients(newRecord)

	// Responder al webhook con un mensaje de éxito
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received successfully"))
}

func sendToClients(record models.Record) {
	events.ClientsLock.Lock()
	defer events.ClientsLock.Unlock()
	for client := range events.Clients {
		// Enviar el registro al cliente a través del canal
		go func(ch chan models.Record) {
			ch <- record
		}(client)
	}
}
