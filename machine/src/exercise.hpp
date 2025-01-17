#include <stdint.h>
#include <string.h>
#include <stdlib.h>
#include <Arduino.h>

#define MAX_DATA_LENGTH 50
#define MAX_EXERCISES 20

typedef struct {
    time_t startedAt;
    time_t finishedAt;
    uint16_t weight;
} ExerciseData;

typedef struct {
    uint64_t userID;
    char originID[64];
    ExerciseData data[MAX_DATA_LENGTH];
    size_t dataLength;

    void addData(ExerciseData data) {
        if (dataLength < MAX_DATA_LENGTH) {
            this->data[dataLength++] = data;
        }
    }
} ExerciseRecord;


ExerciseRecord newExerciseRecord(uint64_t userID, const char* originID) {
    ExerciseRecord record;
    record.userID = userID;
    strncpy(record.originID, originID, sizeof(record.originID));
    return record;
}

ExerciseData newExerciseData(time_t startedAt, time_t finishedAt, uint16_t weight) {
    ExerciseData data;
    data.startedAt = startedAt;
    data.finishedAt = finishedAt;
    data.weight = weight;
    return data;
}

ExerciseRecord copyExerciseRecord(ExerciseRecord record) {
    ExerciseRecord newRecord;
    newRecord.userID = record.userID;
    strncpy(newRecord.originID, record.originID, sizeof(newRecord.originID));
    newRecord.dataLength = record.dataLength;
    for (size_t i = 0; i < record.dataLength; i++) {
        newRecord.data[i] = record.data[i];
    }
    return newRecord;
}

void printExerciseRecords(ExerciseRecord* records, size_t recordLength) {
    for (size_t i = 0; i < recordLength; i++) {
        Serial.printf("Record #%d: %s, %s\n", i, records[i].userID, records[i].originID);
        for (size_t j = 0; j < records[i].dataLength; j++) {
            Serial.printf("  Data #%d: %d, %d, %d\n", j, records[i].data[j].startedAt, records[i].data[j].finishedAt, records[i].data[j].weight);
        }
    }
}