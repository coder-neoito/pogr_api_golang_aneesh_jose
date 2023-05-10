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
	HealthCheck(w http.ResponseWriter, r *http.Request)
	ListAllGames(w http.ResponseWriter, r *http.Request)
	ListGames(w http.ResponseWriter, r *http.Request)
	GetCharacteristics(w http.ResponseWriter, r *http.Request)
	GetFavoriteMap(w http.ResponseWriter, r *http.Request)
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

func (handler profileHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal("success")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func (handler profileHandler) ListAllGames(w http.ResponseWriter, r *http.Request) {

	games, err := handler.service.ListAllGames(r.Context())
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

func (handler profileHandler) GetCharacteristics(w http.ResponseWriter, r *http.Request) {
	var userID string
	var gameCode string

	vars := mux.Vars(r)
	userID = vars["userID"]
	gameCode = vars["gameCode"]

	characteristics, err := handler.service.GetCharacteristics(r.Context(), userID, gameCode)
	if sendErrorResponse(err, w) {
		return
	}

	response := models.Response{
		Data: characteristics,
	}

	res, _ := json.Marshal(response)
	w.Header().Set("content-type", "application/json")
	w.Write(res)
}

func (handler profileHandler) GetFavoriteMap(w http.ResponseWriter, r *http.Request) {
	var userID string
	var gameCode string

	vars := mux.Vars(r)
	userID = vars["userID"]
	gameCode = vars["gameCode"]

	card, err := handler.service.GetFavoriteMap(r.Context(), userID, gameCode)
	if sendErrorResponse(err, w) {
		return
	}

	response := models.Response{
		Data: card,
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
