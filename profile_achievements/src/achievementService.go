package profile_achievements

import (
	"context"
	"errors"
	"fmt"

	"github.com/pogr_api_golang_aneesh_jose/profile_achievements/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type achievementsService struct {
	repository AchievementsRepository
}

type AchievementsService interface {
	GetUserAchievements(ctx context.Context, userID string) ([]models.Achievement, error)
}

func NewAchievementsService(repository AchievementsRepository) AchievementsService {
	return &achievementsService{
		repository: repository,
	}
}

func (service achievementsService) GetUserAchievements(ctx context.Context, userID string) ([]models.Achievement, error) {
	if userID == "" {
		return nil, errors.New("invalid user id")
	}
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return service.repository.GetUserAchievements(ctx, userObjectID)
}
