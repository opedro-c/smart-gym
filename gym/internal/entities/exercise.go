package entities

type Exercise struct {
	Id         uint64
	UserRfId   string `json:"user_rf_id"`
	StartedAt  string `json:"started_at"`
	FinishedAt string `json:"finished_at"`
	Name       string `json:"name"`
}

type Series struct {
	Id          uint64
	StartedAt   string  `json:"started_at"`
	FinishedAt  string  `json:"finished_at"`
	Repetitions uint8   `json:"repetitions"`
	Weight      float32 `json:"weight"`
	ExerciseId  uint64  `json:"exercise_id"`
}
