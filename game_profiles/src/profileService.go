package game_profiles

import (
	"context"

	"github.com/pogr_api_golang_aneesh_jose/game_profiles/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type profileService struct {
	repository ProfileRepository
}

type ProfileService interface {
	ListGames(ctx context.Context, userID string) ([]models.Game, error)
	ListAllGames(ctx context.Context) ([]models.Game, error)
	GetCharacteristics(ctx context.Context, userID string, gameCode string) (models.Characteristics, error)
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

func (service profileService) GetCharacteristics(ctx context.Context, userID string, gameCode string) (models.Characteristics, error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return models.Characteristics{}, err
	}
	return service.repository.GetCharacteristics(ctx, userObjectID, gameCode)
}
