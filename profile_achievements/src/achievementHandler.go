package profile_achievements

import (
	"encoding/json"
	"net/http"

	"github.com/pogr_api_golang_aneesh_jose/profile_achievements/src/models"

	"github.com/gorilla/mux"
)

type achievementsHandler struct {
	service AchievementsService
}

type AchievementsHandler interface {
	HealthCheck(w http.ResponseWriter, r *http.Request)
	GetUserAchievements(w http.ResponseWriter, r *http.Request)
}

func NewAchievementsHandler(service AchievementsService) AchievementsHandler {
	return &achievementsHandler{
		service: service,
	}
}

func (handler achievementsHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal("success")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (handler achievementsHandler) GetUserAchievements(w http.ResponseWriter, r *http.Request) {
	var userID string

	vars := mux.Vars(r)
	userID = vars["userID"]
	achievements, err := handler.service.GetUserAchievements(r.Context(), userID)
	if sendErrorResponse(err, w) {
		return
	}
	response := models.Response{
		Data: achievements,
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
