package machine

import (
	"context"
	"errors"
	"github.com/go-chi/chi/v5"
	"gym-core-service/internal/postgres/connection"
	utils "gym-core-service/pkg"
	s "gym-core-service/pkg/service"
	"net/http"
	"strconv"
)

func machineService(ctx context.Context) *MachineService {
	repository := NewSqlcMachineRepository(connection.GetConnection())
	return NewMachineService(ctx, repository)
}

// @Router			/machines/ [get]
// @Success			200	{object}	any
// @Accept			json
// @Produce			json
func GetMachines(w http.ResponseWriter, r *http.Request) error {
	machines, err := machineService(r.Context()).GetMachines()
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusOK, machines)
	return nil
}

// @Router			/machines/ [get]
// @Success			200	{object}	any
// @Accept			json
// @Produce			json
func CreateMachine(w http.ResponseWriter, r *http.Request) error {
	var input MachineData
	if err := utils.ParseJson(r, &input); err != nil {
		return s.NewServiceError(400, err)
	}

	if err := utils.ValidateJsonStruct(&input); err != nil {
		return s.NewServiceError(400, err)
	}

	machine, err := machineService(r.Context()).CreateMachine(input.Name, input.OriginID)
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusAccepted, machine)
	return nil
}

// @Router			/machines/{id} [get]
// @Success			200	{object}	any
// @Accept			json
// @Produce			json
func DeleteMachine(w http.ResponseWriter, r *http.Request) error {
	machineId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return s.NewServiceError(400, errors.New("machine id invalid"))
	}

	err = machineService(r.Context()).DeleteMachine(int32(machineId))
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusAccepted, nil)
	return nil
}

// @Router			/machines/{id} [get]
// @Success			200	{object}	any
// @Accept			json
// @Produce			json
func UpdateMachine(w http.ResponseWriter, r *http.Request) error {
	machineId, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		return s.NewServiceError(400, errors.New("machine id invalid"))
	}

	var input MachineData
	if err := utils.ParseJson(r, &input); err != nil {
		return s.NewServiceError(400, err)
	}

	if err := utils.ValidateJsonStruct(&input); err != nil {
		return s.NewServiceError(400, err)
	}

	machine, err := machineService(r.Context()).UpdateMachine(int32(machineId), input.Name, input.OriginID)
	if err != nil {
		return err
	}

	utils.WriteJSON(w, http.StatusAccepted, machine)
	return nil
}
