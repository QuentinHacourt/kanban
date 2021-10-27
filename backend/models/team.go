package models

type Team struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type TeamInput struct {
	Name string `json:"name,omitempty"`
}
