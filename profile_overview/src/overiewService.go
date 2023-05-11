package profile_overview

import (
	"context"
	"errors"
	"fmt"

	"github.com/pogr_api_golang_aneesh_jose/profile_overview/src/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type overviewService struct {
	repository OverviewRepository
}

type OverviewService interface {
	GetUser(ctx context.Context, userID string) (models.User, error)
}

func NewOverviewService(repository OverviewRepository) OverviewService {
	return &overviewService{
		repository: repository,
	}
}

func (service overviewService) GetUser(ctx context.Context, userID string) (models.User, error) {
	if userID == "" {
		return models.User{}, errors.New("invalid user id")
	}
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		fmt.Println(err)
		return models.User{}, err
	}
	return service.repository.GetUser(ctx, userObjectID)
}
