package mongo

const (
	DATABASE_NAME = "smart_gym"

	EXERCISES_COLLECTION_NAME = "exercises"
)

type ExerciseCollectionRecord struct {
	UserID     string `json:"user_id" validate:"required"`
	StartedAt  uint64 `json:"started_at" validate:"required"`
	FinishedAt uint64 `json:"finished_at" validate:"required"`
	OriginID   string `json:"origin_id" validate:"required"`
	Data       struct {
		Weight uint32 `json:"weight" validate:"required"`
	} `json:"data" validate:"required,dive"`
}
