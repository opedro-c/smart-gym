package abstract

import (
	"database/sql"
	"fmt"
	"log/slog"
	"reflect"
)

type UseCase[T any, S any] interface {
	Execute(input T) (S, error)
}

type UseCaseTransaction[T any, S any] struct {
	tx      *sql.Tx
	useCase UseCase[T, S]
	input   T
}

func NewUseCaseTransaction[T any, S any](tx *sql.Tx, useCase UseCase[T, S], input T) *UseCaseTransaction[T, S] {
	return &UseCaseTransaction[T, S]{tx, useCase, input}
}

func (a *UseCaseTransaction[T, S]) Execute() (output S, err error) {
	useCaseType := reflect.TypeOf(a.useCase)
	slog.Info(fmt.Sprintf("Executing transaction for use case: %s", useCaseType))
	if output, err = a.useCase.Execute(a.input); err != nil {
		slog.Debug(fmt.Sprintf("Rolling back transaction for use case: %s", useCaseType))
		a.tx.Rollback()
		return output, err
	}
	a.tx.Commit()
	return output, nil
}
