package models

type Achievement struct {
	ID          string   `json:"_id" bson:"_id"`
	Name        string   `json:"name" bson:"name"`
	Provider    string   `json:"provider" bson:"provider"`
	Description string   `json:"description" bson:"description"`
	Thumbnail   string   `json:"thumbnail" bson:"thumbnail"`
	Rewards     []Reward `json:"rewards" bson:"rewards"`
}

type Reward struct {
	ID       string `json:"_id" bson:"_id"`
	Provider string `json:"provider" bson:"provider"`
	Value    int    `json:"value" bson:"value"`
	Type     string `json:"type" bson:"type"`
}
