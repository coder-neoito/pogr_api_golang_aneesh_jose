package game_profiles

import (
	"context"
	"encoding/json"
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
	GetFavoriteMap(ctx context.Context, userID primitive.ObjectID, gameCode string) (models.Card, error)
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

func (repo profileRepository) mapsUserConnectorCollection() *mongo.Collection {
	return repo.client.Database("games_db").Collection("maps_user_connector")
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
	defer curr.Close(ctx)

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
		{"$match": bson.M{
			"$and": bson.A{
				bson.M{"_id": userID},
				bson.M{"games": "gta5"},
			},
		},
		},
		{"$lookup": bson.M{
			"from":         "characteristics",
			"localField":   "games",
			"foreignField": "game_code",
			"as":           "data",
		}},
		{"$unwind": bson.M{"path": "$data"}},
		{"$project": bson.M{
			"data": 1,
			"_id":  0,
		}},
	}

	// Execute the aggregation query
	cursor, err := repo.userCollection().Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	// Iterate over the results and create the response
	var result bson.M

	var data []models.CharacteristicsData
	for cursor.Next(ctx) {
		var res bson.M
		if err := cursor.Decode(&res); err != nil {
			log.Fatal(err)
		}
		result = res
	}

	if result["data"] == nil {
		return models.Characteristics{}, nil
	}
	if result["data"].(primitive.M)["data"] == nil {
		return models.Characteristics{}, nil
	}

	resDecoded, err := json.Marshal(result["data"].(primitive.M)["data"])
	if err != nil {
		fmt.Println(err)
		return models.Characteristics{}, nil
	}

	err = json.Unmarshal(resDecoded, &data)
	if err != nil {
		log.Fatal(err)
		return models.Characteristics{}, nil
	}

	return models.Characteristics{
		Type: "characteristics",
		Name: "Characteristics",
		LeftThumb: models.ThumbData{
			Icon:        "Account",
			IsClickable: false,
		},
		Data: data,
	}, nil
}

func (repo profileRepository) GetFavoriteMap(ctx context.Context, userID primitive.ObjectID, gameCode string) (models.Card, error) {
	pipeline := bson.A{
		bson.M{
			"$match": bson.M{
				"$and": bson.A{
					bson.M{"user_id": userID},
					bson.M{"game_code": gameCode},
				},
			},
		},
		bson.M{
			"$group": bson.M{
				"_id": "$map_id",
				"count": bson.M{
					"$sum": 1,
				},
			},
		},
		bson.M{
			"$sort": bson.M{
				"count": -1,
			},
		},
		bson.M{
			"$limit": 1,
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "maps",
				"localField":   "_id",
				"foreignField": "_id",
				"as":           "map",
			},
		},
		bson.M{
			"$unwind": bson.M{
				"path": "$map",
			},
		},
		bson.M{
			"$lookup": bson.M{
				"from":         "games",
				"localField":   "map.game_code",
				"foreignField": "game_code",
				"as":           "games_details",
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":   1,
				"count": 1,
				"map":   1,
				"game": bson.M{
					"$arrayElemAt": bson.A{
						"$games_details", 0,
					},
				},
			},
		},
		bson.M{
			"$project": bson.M{
				"_id":         0,
				"type":        "card",
				"image":       "$map.url",
				"name":        "$map.name",
				"description": bson.M{"$concat": bson.A{"Most used map for ", "$game.name"}},
				"left_thumb": bson.M{
					"icon": "globe",
				},
			},
		},
		bson.M{
			"$limit": 1,
		},
	}
	fmt.Println(pipeline...)

	// Execute the aggregation query
	cursor, err := repo.mapsUserConnectorCollection().Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)

	var result models.Card
	if cursor.Next(ctx) {
		err := cursor.Decode(&result)
		if err != nil {
			fmt.Println("Error decoding result:", err)
			return result, err
		}
		fmt.Println(result)

	} else {
		fmt.Println("No results found.")
	}
	return result, nil
}
