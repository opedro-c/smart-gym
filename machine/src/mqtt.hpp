#include <exercise.hpp>
#include <ArduinoJson.h>
#include <WiFi.h>
#include <PubSubClient.h>

const char* ssid     = "Wokwi-GUEST";
const char* password = "";

const char* mqtt_server = "maqiatto.com";
const int mqtt_port = 1883;
const char* mqtt_topic = "pedroc_aragao@outlook.com/exercise";
const char* mqtt_user = "pedroc_aragao@outlook.com";
const char* mqtt_password = "senhacomplicadaegrande";

void connectWifi() {
  Serial.print("Connecting to ");
  Serial.println(ssid);
  WiFi.begin(ssid, password);
  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }
  Serial.println("\nWiFi connected.");
}

WiFiClient espClient;
PubSubClient client(espClient);


// Function to serialize ExerciseRecord to JSON
String serializeExerciseRecords(ExerciseRecord* records, size_t recordLength) {
    JsonDocument doc; // Adjust size as needed for larger payloads
    Serial.println("Serializing records to JSON...");
    for (size_t i = 0; i < recordLength; i++) {
        Serial.printf("Record #%d: %s, %s\n", i, records[i].userID, records[i].originID);
        JsonObject recordObj = doc.createNestedObject();
        recordObj["user_id"] = records[i].userID;
        recordObj["origin_id"] = records[i].originID;

        JsonArray dataArray = recordObj.createNestedArray("data");
        for (size_t j = 0; j < records[i].dataLength; j++) {
            Serial.printf("  Data #%d: %d, %d, %d\n", j, records[i].data[j].startedAt, records[i].data[j].finishedAt, records[i].data[j].weight);
            JsonObject dataObj = dataArray.createNestedObject();
            dataObj["started_at"] = records[i].data[j].startedAt;
            dataObj["finished_at"] = records[i].data[j].finishedAt;
            dataObj["weight"] = records[i].data[j].weight;
        }
    }

    // Serialize JSON document to a string
    String output;
    serializeJson(doc, output);
    Serial.println("Serialized JSON:");
    Serial.println(output);
    return output;
}

// Function to publish the ExerciseRecord over MQTT
void publishExerciseRecord(ExerciseRecord* record, size_t recordLength) {
    String payload = serializeExerciseRecords(record, recordLength);
    if (client.publish(mqtt_topic, payload.c_str())) {
        Serial.println("Record published successfully:");
        Serial.println(payload);
    } else {
        Serial.println("Failed to publish record.");
    }
}

// MQTT connection setup
void setupMQTT() {
    client.setServer(mqtt_server, mqtt_port);
    while (!client.connected()) {
        Serial.println("Connecting to MQTT...");
        if (client.connect("ESP32Client", mqtt_user, mqtt_password)) {
            Serial.println("Connected to MQTT broker.");
        } else {
            Serial.print("Failed to connect. State: ");
            Serial.println(client.state());
            delay(2000);
        }
    }
}
