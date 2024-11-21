package exercise

type ExerciseRecord struct {
	UserID   string         `json:"user_id" validate:"required"`
	OriginID string         `json:"origin_id" validate:"required"`
	Data     []ExerciseData `json:"data" validate:"required,dive"`
}

type ExerciseData struct {
	StartedAt  uint64 `json:"started_at" validate:"required"`
	FinishedAt uint64 `json:"finished_at" validate:"required"`
	Weight     uint32 `json:"weight" validate:"required"`
}
