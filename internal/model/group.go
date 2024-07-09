package model

type GroupInfo struct {
	Name string `json:"name,omitempty" validate:"required"`
}

type Group struct {
	GroupInfo
	ID string
}
