#include <lcd.hpp>

const char* ntpServer = "time.google.com";
const long  gmtOffset_sec = 0;
const int   daylightOffset_sec = 3600;

void setup() {
  Serial.begin(9600);
  pinMode(PIN_TRIG, OUTPUT);
  pinMode(PIN_ECHO, INPUT);
  lcd.init();
  lcd.backlight();
  lcd.clear();
  lcd.setCursor(0, 0);
  lcd.print("Initializing...");
  mfrc522.PCD_Init();
  connectWifi();
  setupMQTT();
  configTime(gmtOffset_sec, daylightOffset_sec, ntpServer);
  updateLcdEventGroup = xEventGroupCreate();

  repTimerHandle = xTimerCreate(
    "repTimer",
    pdMS_TO_TICKS(TIMEOUT_CURRENT_SET_MS),
    pdFALSE,
    (void *) 0,
    finishCurrentSet
  );

  setTimerHandle = xTimerCreate(
    "setTimer",
    pdMS_TO_TICKS(TIMEOUT_CURRENT_EXERCISE_MS),
    pdFALSE,
    (void *) 0,
    finishCurrentExercise
  );

  secondsTimerHandle = xTimerCreate(
    "secondsTimer",
    pdMS_TO_TICKS(1000),
    pdTRUE,
    (void *) 0,
    incrementSecond
  );

  xTaskCreate(
    readRFID,
    "readRFID",
    4094,
    NULL,
    0,
    NULL
  );

  xTaskCreate(
    countNumberOfRepetitions,
    "countReps",
    4094,
    NULL,
    2,
    NULL
  );

  xTaskCreate(
    displayStuffOnLCD,
    "displayStuffOnLCD",
    4094,
    NULL,
    1,
    NULL
  );

  lcd.clear();
  displayWaitingRfid();
}

void loop() {
  vTaskDelete(NULL); // remove arduino's main loop
}
