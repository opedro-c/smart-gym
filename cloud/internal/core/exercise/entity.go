package exercise

type ExerciseData struct {
	StartedAt  uint64 `json:"started_at"`
	FinishedAt uint64 `json:"finished_at"`
	Weight     uint32 `json:"weight"`
}

type ExerciseRecord struct {
	UserID   string         `json:"user_id"`
	OriginID string         `json:"origin_id"`
	Data     []ExerciseData `json:"data"`
}
