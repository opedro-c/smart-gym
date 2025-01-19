package sender

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"status-machine-service/internal/core"
)

func HandleConnections(w http.ResponseWriter, r *http.Request) {
	service := core.GetStatusService()

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins, customize this as needed
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading connection: %v", err)
		return
	}
	// defer conn.Close()

	service.AddConnection(conn)
	log.Println("New WebSocket client connected")

	// Clean up on disconnect
	/*
		defer func() {
			service.RemoveConnection(conn)
			log.Println("WebSocket client disconnected")
		}()
	*/
}
