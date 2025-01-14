package machine

import "gym-core-service/internal/postgres/sqlc"

type MachineEntity struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	OriginID string `json:"origin_id"`
}

type MachineData struct {
	Name     string `json:"name" validate:"required,min=3,max=32"`
	OriginID string `json:"origin_id" validate:"required,min=3"`
}

func FromMachineModel(machine sqlc.Machine) MachineEntity {
	return MachineEntity{
		ID:       machine.ID,
		Name:     machine.Name,
		OriginID: machine.OriginID,
	}
}
