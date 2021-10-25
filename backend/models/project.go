package models

type Project struct {
	ID          *int    `json:"id,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}

type ProjectInput struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}
