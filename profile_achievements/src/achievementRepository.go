package profile_achievements

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/pogr_api_golang_aneesh_jose/profile_achievements/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var protocol = GetEnvOrDefault("MONGO_PROTOCOL", "mongodb")
var mongoURI = GetEnvOrDefault("MONGO_URI", "localhost:27017")
var userName = GetEnvOrDefault("MONGO_USERNAME", "root")
var password = GetEnvOrDefault("MONGO_PASSWORD", "password")

type achievementsRepository struct {
	client *mongo.Client
}

type AchievementsRepository interface {
	GetUserAchievements(ctx context.Context, userID primitive.ObjectID) ([]models.Achievement, error)
}

func NewAchievementsRepository() AchievementsRepository {
	connectionURI := fmt.Sprintf("%s://%s:%s@%s/?connect=direct", protocol, userName, password, mongoURI)
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	if err != nil || client == nil {
		panic(errors.New("client is null"))
	}
	return &achievementsRepository{
		client: client,
	}
}

func (repo achievementsRepository) usersCollection() *mongo.Collection {
	return repo.client.Database("games_db").Collection("users")
}

func (repo achievementsRepository) GetUserAchievements(ctx context.Context, userID primitive.ObjectID) ([]models.Achievement, error) {
	var achievements []models.Achievement
	pipeline := bson.A{
		bson.M{
			"$lookup": bson.M{
				"from":         "achievements_user_connector",
				"localField":   "_id",
				"foreignField": "user_id",
				"as":           "achievements",
			},
		},
		bson.M{
			"$project": bson.M{
				"achievements": 1,
			},
		},
		bson.M{
			"$unwind": bson.M{
				"path": "$achievements",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "achievements",
				"localField":   "achievements.achievement_id",
				"foreignField": "_id",
				"as":           "achievement",
			},
		},
		bson.M{
			"$project": bson.M{
				"achievement": 1,
			},
		},
		bson.M{
			"$unwind": bson.M{
				"path": "$achievement",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "rewards",
				"localField":   "achievement.rewards._id",
				"foreignField": "_id",
				"as":           "achievement.rewards",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id": 0,
			},
		},
		bson.M{
			"$replaceRoot": bson.M{
				"newRoot": "$achievement",
			},
		},
	}

	// Execute the aggregation query
	cursor, err := repo.usersCollection().Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	// res := make([]map[string]interface{}, 0)

	// if err := cursor.All(ctx, &res); err != nil {
	if err := cursor.All(ctx, &achievements); err != nil {
		log.Fatal(err)
	}
	// fmt.Println(res)

	return achievements, nil
}
