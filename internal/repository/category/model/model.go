package model

type CategoryInfo struct {
	Name    string `db:"name"`
	GroupId string `db:"group_id"`
}

type Category struct {
	UUID string `db:"id"`
	CategoryInfo
}
