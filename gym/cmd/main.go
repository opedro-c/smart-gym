package main

import (
	"gym/internal/config"
	"gym/internal/mqtt"
	"gym/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	logger.Logger().Println("Starting MQTT client")

	var mosquittoClient mqtt.Client = mqtt.NewMosquittoClient(
		config.MosquittoDomain,
		config.MqttPort,
		config.MqttClientId,
		config.MqttUsername,
		config.MqttPassword,
		config.MqttCleanSession == "false",
	)

	defer mosquittoClient.Disconnect()

	subscriber := mqtt.NewSubscriber(mosquittoClient)

	subscriber.Setup()

	waitForExit := make(chan os.Signal, 1)
	signal.Notify(waitForExit, syscall.SIGINT, syscall.SIGTERM)
	<-waitForExit

	logger.Logger().Println("Exiting MQTT client")
}
