package model

type GroupInfo struct {
	Name string `validate:"required,unique-name" json:"name,omitempty"`
}

type Group struct {
	GroupInfo
	UUID string `json:"id"`
}
