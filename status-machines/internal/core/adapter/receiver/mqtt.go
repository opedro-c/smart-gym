package receiver

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"status-machine-service/internal/core"
)

func ReceiveStatusOnHandler(client MQTT.Client, msg MQTT.Message) {
	service := core.GetStatusService()
	originID := string(msg.Payload())

	status := core.StatusMachine{
		OriginID: originID,
		Status:   true,
	}

	service.SetStatusMachine(status)
}

func ReceiveStatusOffHandler(client MQTT.Client, msg MQTT.Message) {
	service := core.GetStatusService()
	originID := string(msg.Payload())

	status := core.StatusMachine{
		OriginID: originID,
		Status:   false,
	}

	service.SetStatusMachine(status)
}
