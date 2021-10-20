package models

type Story struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"stat"`
}

type StoryInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
