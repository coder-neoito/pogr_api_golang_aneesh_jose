package models

type Card struct {
	Type        *string   `json:"type,default=card"`
	Image       string    `json:"image" bson:"image"`
	Name        string    `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	LeftThumb   ThumbData `json:"left_thumb" bson:"left_thumb"`
}
