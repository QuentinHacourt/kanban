package models

type Story struct {
	ID            *int    `json:"id,omitempty"`
	Title         *string `json:"title,omitempty"`
	Description   *string `json:"description,omitempty"`
	Status        *string `json:"stat,omitempty"`
	Time          *int    `json:"time,omitempty"`
	DeveloperName *string `json:"developer_name,omitempty"`
}

type StoryInput struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	Time          int    `json:"time"`
	DeveloperName string `json:"developer_name,omitempty"`
}
