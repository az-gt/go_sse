package events

import (
	"encoding/json"
	"log"
	"main/internal/models"
	"main/internal/utils"
	"net/http"
	"sync"
	"time"
)

var (
	Clients     = make(map[chan models.Record]bool) // Mapa de canales de clientes conectados
	ClientsLock sync.Mutex                          // Mutex para proteger el mapa de clientes
)

// Función para enviar eventos a los clientes
func SendEvents(w http.ResponseWriter, r *http.Request) {
	// Establecer los headers para indicar que se enviarán eventos
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Obtener el valor del parámetro "plan" de la solicitud
	plan := r.URL.Query().Get("plan")

	// Determinar el tiempo de retraso según el plan del cliente
	delay := utils.GetDelayForPlan(plan)

	// Crear un nuevo canal para este cliente
	clientChan := make(chan models.Record)

	// Registrar el canal del cliente en el mapa
	ClientsLock.Lock()
	Clients[clientChan] = true
	ClientsLock.Unlock()

	// Cerrar el canal del cliente cuando la conexión se cierre
	defer func() {
		ClientsLock.Lock()
		delete(Clients, clientChan)
		ClientsLock.Unlock()
		close(clientChan)
	}()

	// Loop para enviar eventos al cliente con el retraso adecuado
	for record := range clientChan {
		// Esperar el tiempo de retraso antes de enviar el evento
		time.Sleep(delay)

		jsonRecord, err := json.Marshal(record)
		if err != nil {
			log.Println("Error encoding record:", err)
			continue
		}
		// Escribir el evento en la respuesta HTTP
		_, err = w.Write([]byte("data: " + string(jsonRecord) + "\n\n"))
		if err != nil {
			log.Println("Error writing event:", err)
			return
		}
		// Flushear la respuesta para enviar el evento
		w.(http.Flusher).Flush()
	}
}
