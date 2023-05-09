package models

type Characteristics struct {
	Type      string                `json:"type"`
	Name      string                `json:"name"`
	LeftThumb ThumbData             `json:"left_thumb"`
	Data      []CharacteristicsData `json:"data"`
}

type CharacteristicsData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type ThumbData struct {
	Icon        string `json:"image_url"`
	IsClickable bool   `json:"is_clickable"`
}
