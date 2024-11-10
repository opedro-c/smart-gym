package abstract

import (
	"database/sql"
	"fmt"
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
	fmt.Print("Executing transaction for use case: ", reflect.TypeOf(a.useCase))
	if output, err = a.useCase.Execute(a.input); err != nil {
		a.tx.Rollback()
		return output, err
	}
	a.tx.Commit()
	return output, nil
}
