package core

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type StatusService struct {
	usedOriginIDs map[string]bool
	clients       map[*websocket.Conn]bool
	mutex         sync.Mutex
}

var statusServiceInstance *StatusService
var once sync.Once

func GetStatusService() *StatusService {
	once.Do(func() {
		statusServiceInstance = &StatusService{
			usedOriginIDs: make(map[string]bool),
			clients:       make(map[*websocket.Conn]bool),
		}
	})
	return statusServiceInstance
}

func (wss *StatusService) AddConnection(conn *websocket.Conn) {
	wss.mutex.Lock()
	wss.clients[conn] = true
	log.Println("New WebSocket client connected")
	wss.mutex.Unlock()
}

func (wss *StatusService) RemoveConnection(conn *websocket.Conn) {
	wss.mutex.Lock()
	conn.Close()
	delete(wss.clients, conn)
	log.Println("WebSocket client disconnected")
	wss.mutex.Unlock()
}

func (wss *StatusService) SetStatusMachine(status StatusMachine) {
	wss.usedOriginIDs[status.OriginID] = status.Status
	wss.broadcastStatus(status)
}

func (wss *StatusService) broadcastStatus(status StatusMachine) {
	wss.mutex.Lock()
	defer wss.mutex.Unlock()

	message, err := json.Marshal(status)
	if err != nil {
		log.Printf("Error marshaling status to JSON: %v", err)
		return
	}

	for client := range wss.clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Printf("Error sending message: %v", err)
			client.Close()
			delete(wss.clients, client)
		}
	}
}
