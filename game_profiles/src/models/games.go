package models

type Game struct {
	ID   string `json:"id" bson:"_id"`
	Name string `json:"name" bson:"name"`
}
