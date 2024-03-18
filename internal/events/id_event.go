package events

import (
	"encoding/json"
	"log"
	"main/internal/models"
	"net/http"
	"strings"
	"sync"
)

var (
	ClientsID     = make(map[chan models.Record]bool) // Mapa de canales de clientes conectados
	ClientsLockID sync.Mutex                          // Mutex para proteger el mapa de clientes
)

func SendEventsID(w http.ResponseWriter, r *http.Request) {
	// Establecer los headers para indicar que se enviarán eventos
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	ID := strings.TrimPrefix(r.URL.Path, "/events/")

	if ID == "" {
		http.Error(w, "Missing ID parameter", http.StatusBadRequest)
		return
	}
	clientChan := make(chan models.Record)

	ClientsLockID.Lock()
	ClientsID[clientChan] = true
	ClientsLockID.Unlock()

	defer func() {
		ClientsLockID.Lock()
		delete(ClientsID, clientChan)
		ClientsLockID.Unlock()
		close(clientChan)
	}()

	for record := range clientChan {

		jsonRecord, err := json.Marshal(record)
		if err != nil {
			log.Println("Error encoding record:", err)
			continue
		}
		if ID == record.ID { // Escribir el evento en la respuesta HTTP
			_, err = w.Write([]byte("data: " + string(jsonRecord) + "\n\n"))
			if err != nil {
				log.Println("Error writing event:", err)
				return
			}
			// Flushear la respuesta para enviar el evento
			w.(http.Flusher).Flush()
		}
	}
}
