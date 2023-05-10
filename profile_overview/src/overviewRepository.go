package profile_overview

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var protocol = GetEnvOrDefault("MONGO_PROTOCOL", "mongodb")
var mongoURI = GetEnvOrDefault("MONGO_URI", "localhost:27017")
var userName = GetEnvOrDefault("MONGO_USERNAME", "root")
var password = GetEnvOrDefault("MONGO_PASSWORD", "password")

type overviewRepository struct {
	client *mongo.Client
}

type OverviewRepository interface {
	// GetUser(ctx context.Context, userID primitive.ObjectID) (models.User, error)
}

func NewOverviewRepository() OverviewRepository {
	connectionURI := fmt.Sprintf("%s://%s:%s@%s/?connect=direct", protocol, userName, password, mongoURI)
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	if err != nil || client == nil {
		panic(errors.New("client is null"))
	}
	return &overviewRepository{
		client: client,
	}
}

func (repo overviewRepository) gamesCollection() *mongo.Collection {
	return repo.client.Database("games_db").Collection("games")
}
