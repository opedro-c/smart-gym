package abstract

func Zero[T any](value T) T {
	var zero T
	return zero
}
