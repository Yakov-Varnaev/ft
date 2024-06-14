package db

import (
	"database/sql"
	"log/slog"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"

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

type DBObject interface{}

func GetById[ReturnItem DBObject](table string, id uuid.UUID) (*ReturnItem, error) {
	if db == nil {
		panic("Database connection was not initialized yet.")
	}
	var item ReturnItem
	found, err := db.From(table).Where(goqu.C("id").Eq(id)).ScanStruct(&item)
	if err != nil {
		return nil, err
	}
	slog.Info("GetByID", "table", table, "id", id, "found", found)
	if !found {
		return nil, nil
	}
	return &item, nil
}
