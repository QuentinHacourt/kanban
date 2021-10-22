package models

type Developer struct {
	ID   *int    `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type DeveloperInput struct {
	Name *string `json:"name,omitempty"`
}
