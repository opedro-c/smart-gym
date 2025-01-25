#include <MFRC522v2.h>
#include <MFRC522DriverSPI.h>
// #include <MFRC522DriverI2C.h>
#include <MFRC522DriverPinSimple.h>
#include <MFRC522Debug.h>
#include <HTTPClient.h>
#include "distance_sensor.hpp"

#define BIT_WAITING_LIFT_WEIGHT (1 << 2)
#define SDA_PIN 5
// Learn more about using SPI/I2C or check the pin assigment for your board: https://github.com/OSSLibraries/Arduino_MFRC522v2#pin-layout

MFRC522DriverPinSimple ss_pin(SDA_PIN);

MFRC522DriverSPI driver{ss_pin}; // Create SPI driver
// MFRC522DriverI2C driver{};     // Create I2C driver
MFRC522 mfrc522{driver}; // Create MFRC522 instance
char *getUserIdEndpoint = "http://192.168.156.89:8080/rfids/";

uint64_t getUserIdByRfid(byte *rf_id_uid);

void readRFID(void *pvParameters)
{
    byte rf_id_uid[4];
    Serial.println("Reading RFID...");

    while (true)
    {
        if (!mfrc522.PICC_IsNewCardPresent())
        {
            continue;
        }
        Serial.println("Card present");
        // Select one of the cards.
        if (!mfrc522.PICC_ReadCardSerial())
        {
            continue;
        }
        Serial.println("Card read");
        Serial.println(mfrc522.uid.size);
        for (uint8_t i = 0; i < mfrc522.uid.size; i++)
        {
            Serial.print(mfrc522.uid.uidByte[i], HEX);
            rf_id_uid[i] = mfrc522.uid.uidByte[i];
        }
        Serial.println();
        mfrc522.PICC_HaltA();

        uint64_t userId = getUserIdByRfid(rf_id_uid);
        if (userId != -1)
        {
            xQueueSend(userIdQueue, &userId, portMAX_DELAY);
        }
        xEventGroupSetBits(updateLcdEventGroup, BIT_WAITING_LIFT_WEIGHT);
        Serial.println("Publishing machine status...");
        publishMachineStatusOn((char *)origin_id);
    }
}

uint64_t getUserIdByRfid(byte *rf_id_uid)
{
    HTTPClient http;
    uint64_t user_id;

    char url[100];
    snprintf(url, sizeof(url), "%s%02x%02x%02x%02x/user", getUserIdEndpoint,
             rf_id_uid[0], rf_id_uid[1], rf_id_uid[2], rf_id_uid[3]);

    http.begin(url);

    Serial.println("Sending GET request to:");
    Serial.println(url);
    int httpCode = http.GET();

    if (httpCode > 0)
    {
        if (httpCode == HTTP_CODE_OK)
        {
            String payload = http.getString();
            Serial.print("Payload: ");
            Serial.println(payload);
            user_id = strtoull(payload.c_str(), NULL, 10);
            Serial.println(user_id);
            return user_id;
        }
        else
        {
            Serial.printf("[HTTP] GET... failed, error: %s\n", http.errorToString(httpCode).c_str());
            return -1;
        }
    }
    else
    {
        Serial.printf("[HTTP] GET... failed, error: %s\n", http.errorToString(httpCode).c_str());
        return -1;
    }
}
