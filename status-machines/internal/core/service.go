package core

import (
	"encoding/json"
	"fmt"
	"log/slog"
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
	slog.Info("New WebSocket client connected")
	wss.mutex.Unlock()
}

func (wss *StatusService) RemoveConnection(conn *websocket.Conn) {
	wss.mutex.Lock()
	conn.Close()
	delete(wss.clients, conn)
	slog.Info("WebSocket client disconnected")
	wss.mutex.Unlock()
}

func (wss *StatusService) GetLastStatusMachines() []StatusMachine {
	lastStatusMachines := make([]StatusMachine, 0)
	for originID, status := range wss.usedOriginIDs {
		lastStatusMachines = append(lastStatusMachines, StatusMachine{
			OriginID: originID,
			Status:   status,
		})
	}
	return lastStatusMachines
}

func (wss *StatusService) SetStatusMachine(status StatusMachine) {
	wss.usedOriginIDs[status.OriginID] = status.Status
	wss.broadcastStatus(status)
}

func (wss *StatusService) broadcastStatus(status StatusMachine) {
	slog.Info(fmt.Sprintf("Broadcast message of origin_id: [ %s ] with status: [ %t ]", status.OriginID, status.Status))

	wss.mutex.Lock()
	defer wss.mutex.Unlock()

	message, err := json.Marshal(status)
	if err != nil {
		slog.Info(fmt.Sprintf("Error marshaling status to JSON: %s", err))
		return
	}

	for client := range wss.clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			slog.Info(fmt.Sprintf("Error sending message: %s", err))
			client.Close()
			delete(wss.clients, client)
		}
	}
}
