void setup() {
  Serial.begin(115200);
  pinMode(PIN_TRIG, OUTPUT);
  pinMode(PIN_ECHO, INPUT);

  // set up the LCD's number of columns and rows:
  lcd.begin(16, 2);
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
    countNumberOfRepetitions,
    "countReps",
    4094,
    NULL,
    1,
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

  displayWaitingRfid();
}

void loop() {
  delay(10); // this speeds up the simulation, remove later
}
