#include "time.h"
#include "mqtt.hpp"
#include <config.hpp>

#define PIN_TRIG 25
#define PIN_ECHO 32
#define SOUND_SPEED 0.034
#define DISTANCE_THRESHOLD_PUSH 5
#define DISTANCE_THRESHOLD_RELEASE 15
#define TIMEOUT_CURRENT_SET_MS 5000
#define TIMEOUT_CURRENT_EXERCISE_MS 10000 // fixme later
#define FAIL_TO_READ 0

#define BIT_UPDATE_DISPLAY (1 << 0)
#define BIT_WAITING_RFID (1 << 1)

uint8_t numberOfRepetitions = 0;
uint8_t numberOfSets = 0;
uint8_t seconds = 0;
bool isResting = false;

TimerHandle_t repTimerHandle = NULL;
TimerHandle_t setTimerHandle = NULL;
TimerHandle_t secondsTimerHandle = NULL;

EventGroupHandle_t updateLcdEventGroup = NULL;
xQueueHandle userIdQueue = xQueueCreate(1, sizeof(uint64_t));

ExerciseRecord exerciseRecords[MAX_EXERCISES];
ExerciseRecord exerciseRecord = newExerciseRecord(0, origin_id);

void addExerciseRecord()
{
    if (numberOfSets < MAX_EXERCISES)
    {
        exerciseRecords[numberOfSets] = copyExerciseRecord(exerciseRecord);
    }
}

void finishCurrentSet(TimerHandle_t xTimer)
{
    Serial.printf("Set #%d finished: %d reps\n",
                  numberOfSets, numberOfRepetitions);
    addExerciseRecord();
    numberOfRepetitions = 0;
    numberOfSets++;
    seconds = 0;
    isResting = true;
    xEventGroupSetBits(updateLcdEventGroup, BIT_UPDATE_DISPLAY);
}

void finishCurrentExercise(TimerHandle_t xTimer)
{
    Serial.printf("Exercise finished: %d sets\n", numberOfSets);
    publishExerciseRecord(exerciseRecords, numberOfSets);
    numberOfRepetitions = 0;
    numberOfSets = 0;
    seconds = 0;
    xEventGroupSetBits(updateLcdEventGroup, BIT_WAITING_RFID);
    xTimerStop(secondsTimerHandle, portMAX_DELAY);
    publishMachineStatusOff((char *)origin_id);
    exerciseRecord.dataLength = 0;
}

int measureWeightDistance()
{
    // Start a new measurement:
    digitalWrite(PIN_TRIG, LOW);
    vTaskDelay(pdMS_TO_TICKS(0.002));
    digitalWrite(PIN_TRIG, HIGH);
    vTaskDelay(pdMS_TO_TICKS(0.01));
    digitalWrite(PIN_TRIG, LOW);
    int duration = pulseIn(PIN_ECHO, HIGH);
    return duration * SOUND_SPEED / 2;
}

void countNumberOfRepetitions(void *pvParameters)
{
    bool isStillLifted = false;
    bool isStillReleased = true;
    time_t startedAt;
    time_t finishedAt;

    while (true)
    {

        int distance = measureWeightDistance();
        uint64_t newUserId;
        if (xQueueReceive(userIdQueue, &newUserId, 0) == pdTRUE) {
            exerciseRecord.userID = newUserId;
            Serial.print("User ID received: ");
            Serial.println(newUserId);
        }


        if (distance <= DISTANCE_THRESHOLD_PUSH && !isStillLifted && !(distance == FAIL_TO_READ))
        {
            time(&startedAt);
            if (numberOfRepetitions == 1)
            {
                seconds = 0;
                xTimerReset(secondsTimerHandle, portMAX_DELAY);
            }
            isResting = false;
            xTimerReset(repTimerHandle, portMAX_DELAY);
            xTimerReset(setTimerHandle, portMAX_DELAY);
            numberOfRepetitions++;
            xEventGroupSetBits(updateLcdEventGroup, BIT_UPDATE_DISPLAY);
            isStillLifted = true;
            isStillReleased = false;
        }
        else if (distance > DISTANCE_THRESHOLD_RELEASE && !isStillReleased)
        {
            time(&finishedAt);
            exerciseRecord.addData(newExerciseData(startedAt, finishedAt, 50));
            isStillReleased = true;
            isStillLifted = false;
        }
    }
}

void incrementSecond(TimerHandle_t xTimer)
{
    seconds++;
    xEventGroupSetBits(updateLcdEventGroup, BIT_UPDATE_DISPLAY);
}
