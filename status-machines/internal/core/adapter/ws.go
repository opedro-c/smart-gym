package adapter

import (
	"log/slog"
	"net/http"
	"status-machine-service/internal/core"

	"github.com/gorilla/websocket"
)

func HandleConnectionsWsHandler(w http.ResponseWriter, r *http.Request) {
	service := core.GetStatusService()

	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true // Allow all origins
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		slog.Info("Error upgrading connection: ", err)
		return
	}

	service.AddConnection(conn)
	slog.Info("New WebSocket client connected")
}
