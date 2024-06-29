package model

type GroupInfo struct {
	Name string `validator:"required" json:"name"`
}

type Group struct {
	GroupInfo
	UUID string `json:"id"`
}
