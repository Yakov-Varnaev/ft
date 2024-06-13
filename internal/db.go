package db

import (
	"database/sql"
	"log/slog"

	"github.com/doug-martin/goqu/v9"

	_ "github.com/doug-martin/goqu/v9/dialect/sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

var db *goqu.Database

func Init() {
	dbPath := "devdb"
	dialect := goqu.Dialect("sqlite3")
	sqliteDb, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		slog.Error(err.Error())
		panic("Cannot instantiate the db.")
	}
	db = dialect.DB(sqliteDb)
}

const (
	CATEGORY_TABLE  string = "categories"
	GROUPS_TABLE    string = "groups"
	SPENDINGS_TABLE string = "spendings"
)

func GetDB() *goqu.Database {
	return db
}
