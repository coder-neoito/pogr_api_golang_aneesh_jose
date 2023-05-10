package models

type User struct {
	ID    string `json:"id" bson:"_id"`
	Name  string `json:"name" bson:"name"`
	Games []Game `json:"games" bson:"games"`
}
