package models

type Game struct {
	ID        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	SubTitle  string `json:"sub_title" bson:"sub_title"`
	ThumbNail string `json:"thumbnail" bson:"thumbnail"`
	GameCode  string `json:"game_code" bson:"game_code"`
}
