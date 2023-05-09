package game_profiles

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/pogr_api_golang_aneesh_jose/game_profiles/src/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var protocol = GetEnvOrDefault("MONGO_PROTOCOL", "mongodb")
var mongoURI = GetEnvOrDefault("MONGO_URI", "localhost:27017")
var userName = GetEnvOrDefault("MONGO_USERNAME", "root")
var password = GetEnvOrDefault("MONGO_PASSWORD", "password")

type profileRepository struct {
	client *mongo.Client
}

type ProfileRepository interface {
	ListGames(ctx context.Context, userID primitive.ObjectID) ([]models.Game, error)
	ListAllGames(ctx context.Context) ([]models.Game, error)
	GetCharacteristics(ctx context.Context, userID primitive.ObjectID, gameCode string) (models.Characteristics, error)
}

func NewProfileRepository() ProfileRepository {
	connectionURI := fmt.Sprintf("%s://%s:%s@%s/?connect=direct", protocol, userName, password, mongoURI)
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		panic(err)
	}
	if err != nil || client == nil {
		panic(errors.New("client is null"))
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

func (repo profileRepository) ListGames(ctx context.Context, userID primitive.ObjectID) ([]models.Game, error) {

	pipeline := []bson.M{
		{
			"$match": bson.M{"_id": userID},
		},
		{
			"$lookup": bson.M{
				"from":         "games",
				"localField":   "games",
				"foreignField": "game_code",
				"as":           "games",
			},
		},
		{
			"$project": bson.M{
				"_id":   0,
				"games": 1,
			},
		},
	}

	curr, err := repo.userCollection().Aggregate(ctx, pipeline, options.Aggregate())
	// only fetch the first one
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var games []models.Game

	var result []bson.M
	if err := curr.All(ctx, &result); err != nil {
		fmt.Println(err)
		return nil, err
	}
	if gamesArray, ok := result[0]["games"].(primitive.A); ok {
		for _, gameRaw := range gamesArray {
			gameMap := gameRaw.(primitive.M)
			game := models.Game{
				ID:        gameMap["_id"].(primitive.ObjectID).Hex(),
				Name:      gameMap["name"].(string),
				SubTitle:  gameMap["sub_title"].(string),
				ThumbNail: gameMap["thumbnail"].(string),
				GameCode:  gameMap["game_code"].(string),
			}
			games = append(games, game)
		}
	}

	return games, nil

}

func (repo profileRepository) ListAllGames(ctx context.Context) ([]models.Game, error) {
	var games []models.Game

	cursor, err := repo.gamesCollection().Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to list games: %v", err)
	}

	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var game models.Game

		err := cursor.Decode(&game)
		if err != nil {
			return nil, fmt.Errorf("failed to decode game: %v", err)
		}

		games = append(games, game)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("failed to list games: %v", err)
	}

	return games, nil
}

func (repo profileRepository) GetCharacteristics(ctx context.Context, userID primitive.ObjectID, gameCode string) (models.Characteristics, error) {
	pipeline := []bson.M{
		{"$match": bson.M{"userID": userID, "game_code": gameCode}},
		{"$lookup": bson.M{
			"from":         "games",
			"localField":   "game_code",
			"foreignField": "game_code",
			"as":           "game",
		}},
		{"$unwind": bson.M{"path": "$game"}},
		{"$project": bson.M{
			"type":      "characteristics",
			"name":      "$game.name",
			"thumbnail": "$game.thumbnail",
			"data":      "$data",
		}},
	}

	// Execute the aggregation query
	cursor, err := repo.userCollection().Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	// Iterate over the results and create the response
	var results []models.Characteristics
	for cursor.Next(ctx) {
		var result models.Characteristics
		if err := cursor.Decode(&result); err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}

	return models.Characteristics{}, nil
}
