package model

type GroupInfo struct {
	Name string `json:"name,omitempty" validate:"required,unique-name"`
}

type Group struct {
	ID string `json:"id"`
	GroupInfo
}
