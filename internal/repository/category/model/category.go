package model

type Group struct {
	ID   string
	Name string
}

type CategoryInfo struct {
	GroupID string
	Name    string
}

type Category struct {
	ID    string
	Name  string
	Group Group
}
