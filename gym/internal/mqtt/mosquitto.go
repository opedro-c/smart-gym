package mqtt

import (
	"fmt"
	"log/slog"
	"strings"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

type MosquittoClient struct {
	pahoClient       MQTT.Client
	mqttDomain       string
	mqttPort         string
	mqttClientId     string
	mqttUsername     string
	mqttPassword     string
	mqttCleanSession bool
}

func NewMosquittoClient(
	mqttDomain string,
	mqttPort string,
	mqttClientId string,
	mqttUsername string,
	mqttPassword string,
	mqttCleanSession bool,
) *MosquittoClient {

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

	return &MosquittoClient{
		pahoClient,
		mqttDomain,
		mqttPort,
		mqttClientId,
		mqttUsername,
		mqttPassword,
		mqttCleanSession,
	}
}

func (c *MosquittoClient) Subscribe(topic string, handler func(payload []byte)) error {
	slog.Info(fmt.Sprintf("Subscribing to topic: %s\n", topic))
	mosquittoHandler := func(client MQTT.Client, msg MQTT.Message) {
		slog.Info(fmt.Sprintf("\nTopic: %s\nMessage: %s\n", msg.Topic(), msg.Payload()))
		handler(msg.Payload())
	}
	if token := c.pahoClient.Subscribe(topic, 2, mosquittoHandler); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return nil
}

func (c *MosquittoClient) Disconnect() {
	c.pahoClient.Disconnect(250)
}
