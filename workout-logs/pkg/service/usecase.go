package service

type UseCase[T any, S any] interface {
	Execute(input T) (S, error)
}
