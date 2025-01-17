#include <exercise.hpp>
#include <ArduinoJson.h>
#include <WiFi.h>
#include <PubSubClient.h>

const char* ssid     = "POCO X3 Pro";
const char* password = "senhacomplicadaegrande";

const char* mqtt_server = "192.168.95.89";
const int mqtt_port = 1883;
const char* mqtt_exercise_topic = "/exercise";
const char* mqtt_machine_status_topic = "/machine_status";
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
        JsonObject recordObj = doc.createNestedObject();
        recordObj["user_id"] = records[i].userID;
        recordObj["origin_id"] = records[i].originID;

        JsonArray dataArray = recordObj.createNestedArray("data");
        for (size_t j = 0; j < records[i].dataLength; j++) {
            JsonObject dataObj = dataArray.createNestedObject();
            dataObj["started_at"] = records[i].data[j].startedAt;
            dataObj["finished_at"] = records[i].data[j].finishedAt;
            dataObj["weight"] = records[i].data[j].weight;
        }
    }

    // Serialize JSON document to a string
    String output;
    serializeJson(doc, output);
    return output;
}

// Function to publish the ExerciseRecord over MQTT
void publishExerciseRecord(ExerciseRecord* record, size_t recordLength) {
    String payload = serializeExerciseRecords(record, recordLength);
    uint8_t retries = 3;
    while (!client.connected() && retries > 0) {
        Serial.println("Reconnecting to MQTT...");
        if (client.connect("ESP32Client", mqtt_user, mqtt_password)) {
            Serial.println("Connected to MQTT broker.");
        } else {
            Serial.print("Failed to connect. State: ");
            Serial.println(client.state());
            retries--;
        }
    }
    if (client.publish(mqtt_exercise_topic, payload.c_str())) {
        Serial.println("Published message to MQTT.");
    } else {
        Serial.print("Client state: ");
        Serial.println(client.state());
        Serial.println("Failed to publish message to MQTT.");
    }
}

void publishMachineStatus(const char* machine) {
    uint8_t retries = 3;
    while (!client.connected() && retries > 0) {
        Serial.println("Reconnecting to MQTT...");
        if (client.connect("ESP32Client", mqtt_user, mqtt_password)) {
            Serial.println("Connected to MQTT broker.");
        } else {
            Serial.print("Failed to connect. State: ");
            Serial.println(client.state());
            retries--;
        }
    }
    if (client.publish(mqtt_machine_status_topic, machine)) {
        Serial.println("Published message to MQTT.");
    } else {
        Serial.print("Client state: ");
        Serial.println(client.state());
        Serial.println("Failed to publish message to MQTT.");
    }
}

// MQTT connection setup
void setupMQTT() {
    client.setServer(mqtt_server, mqtt_port);
    client.setKeepAlive(60);
    if (client.setBufferSize(4094)) {
        Serial.println("MQTT buffer size set to: 4094");
    } else {
        Serial.println("Failed to set MQTT buffer size.");
    }
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
