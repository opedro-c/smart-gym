package adapter

import (
	"cloud-gym/internal/core/exercise/adapter"
	MQTT "github.com/eclipse/paho.mqtt.golang"
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
		AddBroker(strings.Join([]string{"tcp://", mqttDomain, ":", mqttPort}, "")).
		SetClientID(mqttClientId).
		SetUsername(mqttUsername).
		SetPassword(mqttPassword).
		SetCleanSession(mqttCleanSession)

	pahoClient := MQTT.NewClient(opts)
	if token := pahoClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	if token := pahoClient.Subscribe("/exercise", 2, adapter.CreateExerciseMQTTHandler); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return &pahoClient
}
