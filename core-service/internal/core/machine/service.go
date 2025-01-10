package machine

import (
	"context"
)

type MachineService struct {
	ctx               context.Context
	machineRepository MachineRepository
}

func NewMachineService(ctx context.Context, machineRepository MachineRepository) *MachineService {
	return &MachineService{
		ctx,
		machineRepository,
	}
}

func (s *MachineService) GetMachines() ([]MachineEntity, error) {
	return s.machineRepository.GetMachines(s.ctx)
}

func (s *MachineService) CreateMachine(name string, originID string) (MachineEntity, error) {
	return s.machineRepository.CreateMachine(s.ctx, name, originID)
}

func (s *MachineService) DeleteMachine(id int32) error {
	return s.machineRepository.DeleteMachine(s.ctx, id)
}

func (s *MachineService) UpdateMachine(id int32, name string, originID string) (MachineEntity, error) {
	return s.machineRepository.UpdateMachine(s.ctx, id, name, originID)
}
