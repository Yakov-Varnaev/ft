package model

type GroupInfo struct {
	Name string `db:"name"`
}

type Group struct {
	GroupInfo `db:"group_info"`
	UUID      string `db:"id"`
}
