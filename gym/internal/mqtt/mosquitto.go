package mqtt

import (
	"gym/pkg/logger"
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
		logger.Logger().Fatal(token.Error())
	}
	client := &MosquittoClient{
		pahoClient,
		mqttDomain,
		mqttPort,
		mqttClientId,
		mqttUsername,
		mqttPassword,
		mqttCleanSession,
	}

	return client
}

func (c *MosquittoClient) Subscribe(topic string, handler func(payload []byte)) error {
	logger.Logger().Printf("Subscribing to topic: %s\n", topic)
	mosquittoHandler := func(client MQTT.Client, msg MQTT.Message) {
		logger.Logger().Printf("Topic: %s\nMessage: %s\n", msg.Topic(), msg.Payload())
		handler(msg.Payload())
	}
	if token := c.pahoClient.Subscribe(topic, 0, mosquittoHandler); token.Wait() && token.Error() != nil {
		logger.Logger().Println(token.Error())
		return token.Error()
	}
	return nil
}

func (c *MosquittoClient) Disconnect() {
	c.pahoClient.Disconnect(250)
}
