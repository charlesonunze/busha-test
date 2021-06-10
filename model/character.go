package model

type Character struct {
	Name       string   `json:"name"`
	Height     string   `json:"height"`
	HeightInCm int      `json:"height_in_cm"`
	Gender     string   `json:"gender"`
	Films      []string `json:"films"`
}

type CharacterResponse struct {
	Count    int         `json:"count"`
	Next     string      `json:"next"`
	Previous string      `json:"previous"`
	Results  []Character `json:"results"`
}
