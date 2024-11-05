package database

type Repository interface {
	SaveExercise(exercise Exercise) error
}
