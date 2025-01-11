#define PIN_TRIG    25
#define PIN_ECHO    32
#define SOUND_SPEED 0.034
#define DISTANCE_THRESHOLD 350
#define TIMEOUT_CURRENT_SET_MS 10000
#define TIMEOUT_CURRENT_EXERCISE_MS 120000

uint8_t numberOfRepetitions = 0;
uint8_t numberOfSets = 0;

struct ExercisePayload {
  char* userId;
};

TimerHandle_t repTimerHandle = NULL;
TimerHandle_t setTimerHandle = NULL;

void finishCurrentSet(TimerHandle_t xTimer) {
  Serial.printf("Set #%d finished: %d reps\n",
    numberOfSets, numberOfRepetitions);
  numberOfRepetitions = 0;
  numberOfSets++;
}

void finishCurrentExercise(TimerHandle_t xTimer) {
  Serial.printf("Exercise finished: %d sets\n", numberOfSets);
  numberOfRepetitions = 0;
  numberOfSets = 0;
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
    
    if (distance <= DISTANCE_THRESHOLD && !isStillLifted) {
      xTimerReset(repTimerHandle, portMAX_DELAY);
      xTimerReset(setTimerHandle, portMAX_DELAY);
      numberOfRepetitions++;
      isStillLifted = true;
      Serial.println(numberOfRepetitions);
    } else if (distance > DISTANCE_THRESHOLD) {
      isStillLifted = false;
    }
  }
}

