package adapter

import (
	"cloud-gym/internal/config"
	"cloud-gym/internal/core/exercise/adapter"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"log/slog"
	"strings"
)

func NewMosquittoClient(
	mqttDomain string,
	mqttPort string,
	mqttClientId string,
	mqttUsername string,
	mqttPassword string,
	mqttCleanSession bool,
) *MQTT.Client {

	opts := MQTT.NewClientOptions().
		AddBroker(strings.Join([]string{"mqtt://", mqttDomain, ":", mqttPort}, "")).
		SetClientID(mqttClientId).
		SetUsername(mqttUsername).
		SetPassword(mqttPassword).
		SetCleanSession(mqttCleanSession)

	slog.Info("Connecting to MQTT Broker")
	pahoClient := MQTT.NewClient(opts)
	if token := pahoClient.Connect(); token.Wait() && token.Error() != nil {
		slog.Error("Failed to connect to MQTT Broker", token.Error())
		panic(token.Error())
	}
	topic := strings.Join([]string{config.MqttTopicPrefix, "/exercise"}, "")
	slog.Info("Connected to MQTT Broker")
	slog.Info("Subscribing to MQTT Broker", "topic", topic)

	if token := pahoClient.Subscribe("/exercise", 0, adapter.CreateExerciseMQTTHandler); token.Wait() && token.Error() != nil {
		slog.Error("Failed to subscribe to MQTT Broker", token.Error())
		panic(token.Error())
	}
	return &pahoClient
}
