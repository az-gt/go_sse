package webhooks

import (
	"encoding/json"
	"fmt"
	"log"
	"main/internal/events"
	"main/internal/models"
	"net/http"
	"strings"
)

func HandleWebhookID(w http.ResponseWriter, r *http.Request) {
	// Decodificar los datos del nuevo registro de la solicitud
	var newRecord models.Record
	err := json.NewDecoder(r.Body).Decode(&newRecord)
	if err != nil {
		log.Println("Error decoding webhook payload:", err)
		http.Error(w, "Error decoding webhook payload", http.StatusBadRequest)
		return
	}

	ID := strings.TrimPrefix(r.URL.Path, "/webhook/")
	if ID == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}

	if ID != newRecord.ID {
		fmt.Println(ID, newRecord.ID)
		http.Error(w, "ID parameter does not match record ID", http.StatusBadRequest)
		return
	}

	// Enviar los datos del nuevo registro a los clientes a través de Server-Sent Events
	sendToClientsID(newRecord)

	// Responder al webhook con un mensaje de éxito
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received successfully"))
}

func sendToClientsID(record models.Record) {
	events.ClientsLockID.Lock()
	defer events.ClientsLockID.Unlock()
	for client := range events.ClientsID {
		// Enviar el registro al cliente a través del canal
		go func(ch chan models.Record) {
			ch <- record
		}(client)
	}
}
