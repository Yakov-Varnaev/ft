package model

type CategoryInfo struct {
	Name    string `json:"name" validate:"required,unique-name"`
	GroupId string `json:"group_id" validate:"required,group-exists"`
}

type Category struct {
	UUID string `json:"id,omitempty"`
	CategoryInfo
}
