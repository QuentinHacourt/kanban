package models

type Developer struct {
	ID       *int    `json:"id,omitempty"`
	UserName *string `json:"user_name,omitempty"`
	Password *string `json:"password,omitempty"`
}

type DeveloperInput struct {
	UserName *string `json:"user_name,omitempty"`
	Password *string `json:"password,omitempty"`
}
