#include <stdint.h>
#include <MFRC522v2.h>
#include <MFRC522DriverPinSimple.h>
#include <MFRC522DriverSPI.h>

#define RFID_UID_SIZE 4

struct RfidHandler
{
    byte uid[RFID_UID_SIZE];
    MFRC522 *mfrc522;

    RfidHandler(MFRC522 *mfrc522) : uid(), mfrc522(mfrc522) {}

    ~RfidHandler() {}

    bool readSensor()
    {
        Serial.println("Reading RFID inside handler");
        if (!mfrc522->PICC_IsNewCardPresent())
            return false;
        Serial.println("Card present");
        if (!mfrc522->PICC_ReadCardSerial())
            return false;
        Serial.println("Card read");
        for (int i = 0; i < RFID_UID_SIZE; i++)
        {

            uid[i] = mfrc522->uid.uidByte[i];
        }
        return true;
    }

    byte *getUid()
    {
        return uid;
    }
};
