package game_profiles

import (
	"context"

	"github.com/pogr_api_golang_aneesh_jose/game_profiles/src/models"
)

type profileService struct {
	repository ProfileRepository
}

type ProfileService interface {
	ListGames(ctx context.Context, userID string) ([]models.Game, error)
	ListAllGames(ctx context.Context) ([]models.Game, error)
}

func NewProfileService(repository ProfileRepository) ProfileService {
	return &profileService{
		repository: repository,
	}
}

func (service profileService) ListGames(ctx context.Context, userID string) ([]models.Game, error) {
	return service.repository.ListGames(ctx, userID)
}

func (service profileService) ListAllGames(ctx context.Context) ([]models.Game, error) {
	return service.repository.ListAllGames(ctx)
}
