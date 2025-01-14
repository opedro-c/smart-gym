package machine

import (
	"context"
	"gym-core-service/internal/postgres/sqlc"
)

type MachineRepository interface {
	GetMachines(ctx context.Context) ([]MachineEntity, error)
	CreateMachine(ctx context.Context, name string, originID string) (MachineEntity, error)
	DeleteMachine(ctx context.Context, id int32) error
	UpdateMachine(ctx context.Context, id int32, name string, originID string) (MachineEntity, error)
}

type SqlcMachineRepository struct {
	db *sqlc.Queries
}

func NewSqlcMachineRepository(db *sqlc.Queries) MachineRepository {
	return &SqlcMachineRepository{
		db: db,
	}
}

func (r *SqlcMachineRepository) GetMachines(ctx context.Context) ([]MachineEntity, error) {
	machines, err := r.db.GetMachines(ctx)
	if err != nil {
		return nil, err
	}

	machineEntities := make([]MachineEntity, len(machines))
	for i, machine := range machines {
		machineEntities[i] = MachineEntity{
			ID:       machine.ID,
			Name:     machine.Name,
			OriginID: machine.OriginID,
		}
	}

	return machineEntities, nil
}

func (r *SqlcMachineRepository) CreateMachine(ctx context.Context, name string, originID string) (MachineEntity, error) {
	model, err := r.db.CreateMachine(ctx, sqlc.CreateMachineParams{
		Name:     name,
		OriginID: originID,
	})

	return FromMachineModel(model), err
}

func (r *SqlcMachineRepository) DeleteMachine(ctx context.Context, id int32) error {
	return r.db.DeleteMachine(ctx, id)
}

func (r *SqlcMachineRepository) UpdateMachine(ctx context.Context, id int32, name string, originID string) (MachineEntity, error) {
	machine, err := r.db.UpdateMachine(ctx, sqlc.UpdateMachineParams{
		ID:       id,
		Name:     name,
		OriginID: originID,
	})

	return FromMachineModel(machine), err

}
