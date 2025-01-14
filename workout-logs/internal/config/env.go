package config

import (
	"os"
)

var (
	MosquittoDomain  = os.Getenv("MOSQUITTO_DOMAIN")
	MqttPort         = os.Getenv("MQTT_PORT")
	MqttClientId     = os.Getenv("MQTT_CLIENT_ID")
	MqttUsername     = os.Getenv("MQTT_USERNAME")
	MqttPassword     = os.Getenv("MQTT_PASSWORD")
	MqttCleanSession = os.Getenv("MQTT_CLEAN_SESSION")
	MongoUrl         = os.Getenv("MONGODB_URL")
)
