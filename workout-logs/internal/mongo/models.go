package mongo

import "time"

const (
	DATABASE_NAME = "smart_gym"

	EXERCISES_COLLECTION_NAME = "exercises"
)

type ExerciseCollectionRecord struct {
	UserID     string    `json:"user_id" bson:"user_id" validate:"required"`
	StartedAt  time.Time `json:"started_at" bson:"started_at" validate:"required"`
	FinishedAt time.Time `json:"finished_at" bson:"finished_at" validate:"required"`
	OriginID   string    `json:"origin_id" bson:"origin_id" validate:"required"`
	Data       struct {
		Weight uint64 `json:"weight" bson:"weight" validate:"required"`
	} `json:"data" bson:"data" validate:"required,dive"`
}
