package abstract

import "database/sql"

type UseCase[T any] interface {
	Execute(input T) error
}

type AbstractUseCase[T any] struct {
	UseCase[T]
	Tx *sql.Tx

}

func (a *AbstractUseCase[T]) Execute(input T) error {
	if err := a.UseCase.Execute(input); err != nil {
		a.Tx.Rollback()
		return err
	}
	a.Tx.Commit()
	return nil
}
