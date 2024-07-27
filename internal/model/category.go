package model

type CategoryInfo struct {
	GroupID string `json:"group_id" validate:"required"`
	Name    string `json:"name" validate:"required,unique-name"`
}

type Category struct {
	ID    string `json:"id"`
	Name  string `json:"name" validate:"required,unique-name"`
	Group Group  `json:"group"`
}
