package profile_overview

import (
	"encoding/json"
	"net/http"
)

type overviewHandler struct {
	service OverviewService
}

type OverviewHandler interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

func NewOverviewHandler(service OverviewService) OverviewHandler {
	return &overviewHandler{
		service: service,
	}
}

func (handler overviewHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal("success")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
