#include <MFRC522v2.h>
#include <MFRC522DriverSPI.h>
#include <MFRC522DriverI2C.h>
#include <MFRC522DriverPinSimple.h>
#include <MFRC522Debug.h>
#include <rfid.hpp>


void setup()
{
  Serial.begin(9600);
  // while (!Serial); 
  // MFRC522DriverPinSimple ss_pin(5);

  // MFRC522DriverSPI driver{ss_pin}; 
  // MFRC522 mfrc522{driver};         
  // mfrc522.PCD_Init();              
  // RfidHandler rfidHandler{&mfrc522};

  // while (true)
  // {
  //   Serial.println("Reading RFID");
  //   if (rfidHandler.readSensor())
  //   {
  //     Serial.print("Card UID: ");
  //     for (int i = 0; i < RFID_UID_SIZE; i++)
  //     {
  //       Serial.print(rfidHandler.getUid()[i], HEX);
  //       Serial.print(" ");
  //     }
  //     Serial.println();
  //   }
  //   delay(1000);
  // }
  

  // xTaskCreate([](void *pvParameters) {
  //   RfidHandler *rfidHandler = (RfidHandler *)pvParameters;
  //   Serial.println("RFID Task Started");
  //   while (true)
  //   {
  //     Serial.println("Reading RFID");
  //     if (rfidHandler->readSensor())
  //     {
  //       Serial.print("Card UID: ");
  //       for (int i = 0; i < RFID_UID_SIZE; i++)
  //       {
  //         Serial.print(rfidHandler->getUid()[i], HEX);
  //         Serial.print(" ");
  //       }
  //       Serial.println();
  //     }
  //     vTaskDelay(1000 / portTICK_PERIOD_MS);
  //   }

  //   vTaskDelete(NULL);
  // }, "rfidTask", 4096, &rfidHandler, 1, NULL);
}

void loop()
{
  Serial.println("Hello World");
}