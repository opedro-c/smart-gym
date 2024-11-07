package entities

type Exercise struct {
	Id          uint64
	UserRfId    string
	StartedAt  string
	FinishedAt string
	Name        string
}

type Series struct {
	Id          uint64
	StartedAt  string
	FinishedAt string
	Repetitions uint8
	Weight      float32
	ExerciseId  uint64
}
