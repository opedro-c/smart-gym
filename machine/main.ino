void setup() {
  Serial.begin(115200);
  Serial.println("Hello, ESP32!");
  pinMode(PIN_TRIG, OUTPUT);
  pinMode(PIN_ECHO, INPUT);

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

  xTaskCreate(
    countNumberOfRepetitions,
    "countReps",
    4094,
    NULL,
    1,
    NULL
  );
}

void loop() {
  delay(10); // this speeds up the simulation, remove later
}
