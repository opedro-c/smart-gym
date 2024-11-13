package main

import (
	"gym/internal/config"
	"gym/internal/mqtt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	slog.Info("Starting MQTT client")

	var mosquittoClient mqtt.Client = mqtt.NewMosquittoClient(
		config.MosquittoDomain,
		config.MqttPort,
		config.MqttClientId,
		config.MqttUsername,
		config.MqttPassword,
		config.MqttCleanSession == "true",
	)

	defer mosquittoClient.Disconnect()

	subscriber := mqtt.NewSubscriber(mosquittoClient)

	subscriber.Setup()

	waitForExit := make(chan os.Signal, 1)
	signal.Notify(waitForExit, syscall.SIGINT, syscall.SIGTERM)
	<-waitForExit

	slog.Info("Exiting MQTT client")
}
