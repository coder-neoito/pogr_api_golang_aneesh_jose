package game_profiles

import (
	"encoding/json"
	"net/http"

	"github.com/pogr_api_golang_aneesh_jose/game_profiles/src/models"
)

type profileHandler struct {
	service ProfileService
}

type ProfileHandler interface {
	ListGames(w http.ResponseWriter, r *http.Request)
}

func NewProfilehandler(service ProfileService) ProfileHandler {
	return &profileHandler{
		service: service,
	}
}

func (handler profileHandler) ListGames(w http.ResponseWriter, r *http.Request) {
	var userID string

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
