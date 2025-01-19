package adapter

import (
	"net/http"
	"status-machine-service/internal/core"
	"status-machine-service/pkg"
)

func GetLastStatusHttpHandler(w http.ResponseWriter, r *http.Request) {
	service := core.GetStatusService()
	lastStatus := service.GetLastStatusMachines()
	pkg.WriteJSON(w, http.StatusOK, lastStatus)
}
