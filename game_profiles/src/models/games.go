package models

import (
	"encoding/json"
	"io"
)

type GameParser struct{}

func NewGameParser() *GameParser {
	return &GameParser{}
}

func (p *GameParser) Parse(body io.Reader) (*Game, error) {
	var game Game

	err := json.NewDecoder(body).Decode(&game)
	if err != nil {
		return nil, err
	}

	return &game, nil
}

type Game struct {
	ID        string `json:"id" bson:"_id"`
	Name      string `json:"name" bson:"name"`
	SubTitle  string `json:"sub_title" bson:"sub_title"`
	ThumbNail string `json:"thumbnail" bson:"thumbnail"`
}
