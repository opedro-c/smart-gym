// include the library code:
#include "distance_sensor.hpp"
#include <LiquidCrystal.h>

// initialize the library with the numbers of the interface pins
LiquidCrystal lcd(19, 23, 18, 17, 16, 15);

void displayRepsAndSets();
void displayTime();
void displayResting();
void displayWorkOut();
void displayWaitingRfid();

void displayStuffOnLCD(void *pvParameters) {

  while (true) {
    EventBits_t uxBits = xEventGroupWaitBits(
                           updateLcdEventGroup,   /* The event group being tested. */
                           BIT_UPDATE_DISPLAY | BIT_WAITING_RFID, /* The bits within the event group to wait for. */
                           pdTRUE,        /* bits cleared before returning. */
                           pdFALSE,       /* Don't wait for both bits, either bit will do. */
                           portMAX_DELAY
                         );
    lcd.clear();
    if ( ( uxBits & BIT_UPDATE_DISPLAY ) != 0 ) {
      displayRepsAndSets();
      if (isResting) {
        displayResting();
      } else {
        displayWorkOut();
      }
      displayTime();
    } else if ( ( uxBits & BIT_WAITING_RFID ) != 0 ) {
      displayWaitingRfid();
    }
  }
}

void displayRepsAndSets() {
  lcd.setCursor(0, 0);
  lcd.print("Sets:");
  lcd.print(numberOfSets);
  lcd.print("|");
  lcd.print("Reps:");
  lcd.print(numberOfRepetitions);
}

void displayTime() {
  lcd.setCursor(11, 1);
  uint8_t minutes = seconds / 60;
  lcd.print("0");
  lcd.print(minutes);
  lcd.print(":");
  if (seconds < 10) {
    lcd.print("0");
  }
  lcd.print(seconds);
}

void displayResting() {
  lcd.setCursor(0, 1);
  lcd.print("Resting:");
}

void displayWorkOut() {
  lcd.setCursor(0, 1);
  lcd.print("Work out:");
}

void displayWaitingRfid() {
  lcd.setCursor(0, 0);
  lcd.print("Approach card");
  lcd.setCursor(0, 1);
  lcd.print("to start");
}

