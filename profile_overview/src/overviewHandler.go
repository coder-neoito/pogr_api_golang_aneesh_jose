package profile_overview

import (
	"encoding/json"
	"net/http"

	"github.com/pogr_api_golang_aneesh_jose/profile_overview/src/models"

	"github.com/gorilla/mux"
)

type overviewHandler struct {
	service OverviewService
}

type OverviewHandler interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
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

func (handler overviewHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	var userID string

	vars := mux.Vars(r)
	userID = vars["userID"]
	games, err := handler.service.GetUser(r.Context(), userID)
	if sendErrorResponse(err, w) {
		return
	}
	response := models.Response{
		Data: games,
	}

	res, _ := json.Marshal(response)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func sendErrorResponse(err error, w http.ResponseWriter) bool {
	if err == nil {
		return false
	}
	response := models.Response{
		Data: err.Error(),
		Err:  true,
	}
	res, _ := json.Marshal(response)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
	w.WriteHeader(http.StatusBadRequest)
	return true
}
