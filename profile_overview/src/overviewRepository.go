package profile_overview

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/pogr_api_golang_aneesh_jose/profile_overview/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	GetUser(ctx context.Context, userID primitive.ObjectID) (models.User, error)
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

func (repo overviewRepository) usersCollection() *mongo.Collection {
	return repo.client.Database("games_db").Collection("users")
}

func (repo overviewRepository) GetUser(ctx context.Context, userID primitive.ObjectID) (models.User, error) {
	var user models.User
	pipeline := bson.A{
		bson.M{
			"$match": bson.M{"_id": userID},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "games",
				"localField":   "games",
				"foreignField": "game_code",
				"as":           "games",
			},
		},
		bson.M{
			"$limit": 1,
		},
	}

	// Execute the aggregation query
	cursor, err := repo.usersCollection().Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		err := cursor.Decode(&user)
		if err != nil {
			fmt.Println("Error decoding user:", err)
			return user, err
		}

	} else {
		fmt.Println("No users found.")
	}
	return user, nil
}
