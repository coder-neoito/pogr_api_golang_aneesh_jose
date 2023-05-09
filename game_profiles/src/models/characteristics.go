package models

type Characteristics struct {
	Type      string              `json:"type"`
	Name      string              `json:"name"`
	Thumbnail string              `json:"thumbnail"`
	Data      CharacteristicsData `json:"data"`
}

type CharacteristicsData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
