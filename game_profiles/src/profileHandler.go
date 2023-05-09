package game_profiles

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pogr_api_golang_aneesh_jose/game_profiles/src/models"
)

type profileHandler struct {
	service ProfileService
}

type ProfileHandler interface {
	ListGames(w http.ResponseWriter, r *http.Request)
	ListAllGames(w http.ResponseWriter, r *http.Request)
	HealthCheck(w http.ResponseWriter, r *http.Request)
}

func NewProfilehandler(service ProfileService) ProfileHandler {
	return &profileHandler{
		service: service,
	}
}

func (handler profileHandler) ListGames(w http.ResponseWriter, r *http.Request) {
	var userID string

	vars := mux.Vars(r)
	userID = vars["userID"]
	games, err := handler.service.ListGames(r.Context(), userID)
	var response models.Response
	if err != nil {
		response = models.Response{
			Err:  true,
			Data: err.Error(),
		}
	}
	response = models.Response{
		Data: games,
	}

	res, _ := json.Marshal(response)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func (handler profileHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal("success")
	w.Write(res)
	w.WriteHeader(http.StatusOK)
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

func (handler profileHandler) ListAllGames(w http.ResponseWriter, r *http.Request) {

	games, err := handler.service.ListAllGames(r.Context())
	var response models.Response
	if err != nil {
		response = models.Response{
			Err:  true,
			Data: err.Error(),
		}
	}
	response = models.Response{
		Data: games,
	}

	res, _ := json.Marshal(response)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}
