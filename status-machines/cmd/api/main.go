package main

import (
	"log/slog"
	"net/http"
	"os"
	"status-machine-service/internal/config"
	"status-machine-service/internal/core/adapter"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	setUpLogger()

	/// ----------
	/// MQTT
	/// ----------
	opts := MQTT.NewClientOptions().
		AddBroker(strings.Join([]string{"mqtt://", config.MosquittoDomain, ":", config.MqttPort}, "")).
		SetClientID(config.MqttClientId).
		SetUsername(config.MqttUsername).
		SetPassword(config.MqttPassword).
		SetCleanSession(config.MqttCleanSession == "true")

	slog.Info("Connecting to MQTT Broker")

	mqttClient := MQTT.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		slog.Error("Failed to connect to MQTT Broker", token.Error())
		panic(token.Error())
	}
	defer mqttClient.Disconnect(1000)

	if token := mqttClient.Subscribe("/machine_status/on", 0, adapter.ReceiveStatusOnMqttHandler); token.Wait() && token.Error() != nil {
		slog.Error("Failed to subscribe to MQTT Broker", token.Error())
		panic(token.Error())
	}
	if token := mqttClient.Subscribe("/machine_status/off", 0, adapter.ReceiveStatusOffMqttHandler); token.Wait() && token.Error() != nil {
		slog.Error("Failed to subscribe to MQTT Broker", token.Error())
		panic(token.Error())
	}
	slog.Info("Connected to MQTT Broker")

	/// ----------
	/// WS & HTTP
	/// ----------
	http.HandleFunc("/ws", withCORS(adapter.HandleConnectionsWsHandler))
	http.HandleFunc("/status", withCORS(adapter.GetLastStatusHttpHandler))

	go func() {
		slog.Info("WebSocket server started on :7070")
	}()
	http.ListenAndServe(":7070", nil)
}

// CORS middleware function
func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins, modify as needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight request (OPTIONS)
		if r.Method == http.MethodOptions {
			return
		}

		// Call the next handler
		next(w, r)
	}

func setUpLogger() {
	logLevelEnv := config.LogLevel
	var logLevel slog.Level
	switch logLevelEnv {
	case "DEBUG":
		slog.Info("Setting log level to DEBUG")
		logLevel = slog.LevelDebug
	case "INFO":
		slog.Info("Setting log level to INFO")
		logLevel = slog.LevelInfo
	case "ERROR":
		slog.Info("Setting log level to ERROR")
		logLevel = slog.LevelError
	}
	opts := &slog.HandlerOptions{
		Level: logLevel,
	}
	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
