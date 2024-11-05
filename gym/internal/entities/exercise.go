package entities

type Exercise struct {
	Id          uint64
	UserRfId    string
	Started_at  string
	Finished_at string
	Name        string
}

type Series struct {
	Id          uint64
	Started_at  string
	Finished_at string
	Repetitions uint64
	Weight      float64
	ExerciseId  uint64
}
