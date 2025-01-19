package adapter

import (
	"status-machine-service/internal/core"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func ReceiveStatusOnMqttHandler(client MQTT.Client, msg MQTT.Message) {
	service := core.GetStatusService()
	originID := string(msg.Payload())

	status := core.StatusMachine{
		OriginID: originID,
		Status:   true,
	}

	service.SetStatusMachine(status)
}

func ReceiveStatusOffMqttHandler(client MQTT.Client, msg MQTT.Message) {
	service := core.GetStatusService()
	originID := string(msg.Payload())

	status := core.StatusMachine{
		OriginID: originID,
		Status:   false,
	}

	service.SetStatusMachine(status)
}
