package exercise

import (
	m "cloud-gym/internal/mongo"
	"time"
)

type ExerciseRecord struct {
	UserID   uint64         `json:"user_id" validate:"required"`
	OriginID string         `json:"origin_id" validate:"required"`
	Data     []ExerciseData `json:"data" validate:"required,dive"`
}

type ExerciseData struct {
	StartedAt  uint64 `json:"started_at" validate:"required"`
	FinishedAt uint64 `json:"finished_at" validate:"required"`
	Weight     uint64 `json:"weight" validate:"required"`
}

func NewExerciseCollectionRecord(exerciseRecords []ExerciseRecord) []m.ExerciseCollectionRecord {
	// Estimate the size of the result slice for better memory allocation
	totalRecords := 0
	for _, record := range exerciseRecords {
		totalRecords += len(record.Data)
	}

	result := make([]m.ExerciseCollectionRecord, 0, totalRecords)

	for _, record := range exerciseRecords {
		for _, data := range record.Data {
			result = append(result, m.ExerciseCollectionRecord{
				UserID:     record.UserID,
				OriginID:   record.OriginID,
				StartedAt:  time.UnixMilli(int64(data.StartedAt)).UTC(),
				FinishedAt: time.UnixMilli(int64(data.StartedAt)).UTC(),
				Data: struct {
					Weight uint64 `json:"weight" bson:"weight" validate:"required"`
				}{
					Weight: data.Weight,
				},
			})
		}
	}

	return result
}
