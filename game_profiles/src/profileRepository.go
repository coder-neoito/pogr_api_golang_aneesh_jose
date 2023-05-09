package game_profiles

import (
	"context"

	"github.com/pogr_api_golang_aneesh_jose/game_profiles/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type profileRepository struct {
	client *mongo.Client
}

type ProfileRepository interface {
	ListGames(ctx context.Context, userID string) ([]models.Game, error)
}

func NewProfileRepository() ProfileRepository {
	clientOptions := options.Client().ApplyURI("mongodb://root:password@localhost:27017")
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		return nil
	}
	if err != nil || client == nil {
		panic(err)
	}
	return &profileRepository{
		client: client,
	}
}

func (repo profileRepository) gamesCollection() *mongo.Collection {
	return repo.client.Database("games_db").Collection("games")
}

func (repo profileRepository) userCollection() *mongo.Collection {
	return repo.client.Database("games_db").Collection("users")
}

func (repo profileRepository) ListGames(ctx context.Context, userID string) ([]models.Game, error) {

	filter := bson.M{"userID": userID}
	projection := bson.M{"games": 1}
	var result bson.M

	err := repo.userCollection().FindOne(ctx, filter, options.FindOne().SetProjection(projection)).Decode(&result)
	if err != nil {
		return nil, err
	}
	var ok bool
	games, ok := result["games"].([]models.Game)
	if !ok {
		return nil, err
	}

	return games, nil
}
