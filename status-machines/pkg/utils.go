package pkg

import (
	"bytes"
	"encoding/json"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"net/http"
)

func ParseJson(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

func ParseMQTTPayload(message MQTT.Message, v interface{}) error {
	return json.NewDecoder(bytes.NewReader(message.Payload())).Decode(v)
}

func WriteJSON(w http.ResponseWriter, status int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}
