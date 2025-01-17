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
char *getUserIdEndpoint = "http://192.168.95.89:8080/rfids/";

void getUserIdByRfid(byte *rf_id_uid);

void readRFID(void *pvParameters)
{
    byte rf_id_uid[4];

    while (true)
    {
        if (!mfrc522.PICC_IsNewCardPresent())
        {
            continue;
        }

        // Select one of the cards.
        if (!mfrc522.PICC_ReadCardSerial())
        {
            continue;
        }
        Serial.println(mfrc522.uid.size);
        for (uint8_t i = 0; i < mfrc522.uid.size; i++)
        {
            Serial.print(mfrc522.uid.uidByte[i], HEX);
            rf_id_uid[i] = mfrc522.uid.uidByte[i];
        }
        Serial.println();
        mfrc522.PICC_HaltA();

        getUserIdByRfid(rf_id_uid);

        xEventGroupSetBits(updateLcdEventGroup, BIT_WAITING_LIFT_WEIGHT);
        Serial.println("Publishing machine status...");
        publishMachineStatus((char *)rf_id_uid);
    }
}

void getUserIdByRfid(byte *rf_id_uid)
{
    HTTPClient http;

    char url[100];
    snprintf(url, sizeof(url), "%s%02x%02x%02x%02x/user", getUserIdEndpoint,
             rf_id_uid[0], rf_id_uid[1], rf_id_uid[2], rf_id_uid[3]);

    http.begin(url);

    int httpCode = http.GET();

    if (httpCode > 0)
    {
        if (httpCode == HTTP_CODE_OK)
        {
            String payload = http.getString();
            Serial.println(payload);
            user_id = strtoull(payload.c_str(), NULL, 10);
            Serial.println(user_id);
        }
    }
    else
    {
        Serial.printf("[HTTP] GET... failed, error: %s\n", http.errorToString(httpCode).c_str());
    }
}
