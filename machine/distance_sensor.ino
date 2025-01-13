#define PIN_TRIG                    25
#define PIN_ECHO                    32
#define SOUND_SPEED                 0.034
#define DISTANCE_THRESHOLD_PUSH     350
#define DISTANCE_THRESHOLD_RELEASE  380
#define TIMEOUT_CURRENT_SET_MS      10000
#define TIMEOUT_CURRENT_EXERCISE_MS 30000 //fixme later

#define BIT_UPDATE_DISPLAY  ( 1 << 0 )
#define BIT_WAITING_RFID    ( 1 << 1 )

uint8_t numberOfRepetitions = 0;
uint8_t numberOfSets        = 0;
uint8_t seconds             = 0;
bool    isResting           = false;

TimerHandle_t repTimerHandle     = NULL;
TimerHandle_t setTimerHandle     = NULL;
TimerHandle_t secondsTimerHandle = NULL;

EventGroupHandle_t updateLcdEventGroup = NULL;

void finishCurrentSet(TimerHandle_t xTimer) {
  Serial.printf("Set #%d finished: %d reps\n",
                numberOfSets, numberOfRepetitions);
  numberOfRepetitions = 0;
  numberOfSets++;
  seconds = 0;
  isResting = true;
  xEventGroupSetBits(updateLcdEventGroup, BIT_UPDATE_DISPLAY);
}

void finishCurrentExercise(TimerHandle_t xTimer) {
  Serial.printf("Exercise finished: %d sets\n", numberOfSets);
  numberOfRepetitions = 0;
  numberOfSets = 0;
  seconds = 0;
  xEventGroupSetBits(updateLcdEventGroup, BIT_WAITING_RFID);
  xTimerStop(secondsTimerHandle, portMAX_DELAY);
}

int measureWeightDistance() {
  // Start a new measurement:
  digitalWrite(PIN_TRIG, HIGH);
  digitalWrite(PIN_TRIG, LOW);
  // Read the result:
  int duration = pulseIn(PIN_ECHO, HIGH);
  return duration * SOUND_SPEED / 2;
}

void countNumberOfRepetitions(void *pvParameters) {
  bool isStillLifted = false;

  while (true) {

    int distance = measureWeightDistance();

    if (distance <= DISTANCE_THRESHOLD_PUSH
        && !isStillLifted) {
      if (numberOfRepetitions == 1) {
        Serial.println("Starting seconds timer");
        xTimerReset(secondsTimerHandle, portMAX_DELAY);
      }
      isResting = false;
      xTimerReset(repTimerHandle, portMAX_DELAY);
      xTimerReset(setTimerHandle, portMAX_DELAY);
      numberOfRepetitions++;
      xEventGroupSetBits(updateLcdEventGroup, BIT_UPDATE_DISPLAY);
      isStillLifted = true;
    } else if (distance > DISTANCE_THRESHOLD_RELEASE) {
      isStillLifted = false;
    }
  }
}

void incrementSecond(TimerHandle_t xTimer) {
  seconds++;
  Serial.println(seconds);
  xEventGroupSetBits(updateLcdEventGroup, BIT_UPDATE_DISPLAY);
}

