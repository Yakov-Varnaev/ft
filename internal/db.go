package db

import (
	"database/sql"
	"fmt"
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
type WriteDBObject interface{}

type QueryProcessor interface {
	Process(query *goqu.SelectDataset) *goqu.SelectDataset
}

func Create[WriteData WriteDBObject, ReturnItem DBObject](table string, data *WriteData) (*ReturnItem, error) {
	var item ReturnItem
	_, err := db.Insert(table).Rows(data).Executor().Exec()

	if err != nil {
		return nil, err
	}
	return &item, nil
}

func GetById[ReturnItem DBObject](table string, id uuid.UUID, p QueryProcessor) (*ReturnItem, error) {
	if db == nil {
		panic("Database connection was not initialized yet.")
	}
	var item ReturnItem
	query := db.From(table).Where(goqu.I(table + "." + "id").Eq(id))
	if p != nil {
		query = p.Process(query)
	}
	found, err := query.ScanStruct(&item)
	if err != nil {
		return nil, err
	}
	slog.Info("GetByID", "table", table, "id", id, "found", found)
	if !found {
		return nil, nil
	}
	return &item, nil
}

func DeleteById(table string, id uuid.UUID) (*uuid.UUID, error) {
	if db == nil {
		panic("db is fucked")
	}
	_, err := db.Delete(table).Where(goqu.C("id").Eq(id)).Executor().Exec()
	if err != nil {
		return nil, err
	}
	return &id, nil
}

func List[ReturnItem DBObject](table string, processor QueryProcessor) (
	*[]ReturnItem, error,
) {
	items := []ReturnItem{}
	query := db.From(table)
	if processor != nil {
		query = processor.Process(query)
	}

	err := query.ScanStructs(&items)
	if err != nil {
		return nil, err
	}
	return &items, nil
}

func Update[WriteData WriteDBObject, ReturnItem DBObject](table string, id uuid.UUID, data *WriteData) (*ReturnItem, error) {
	_, err := db.Update(table).Where(goqu.C("id").Eq(id)).Set(&data).Executor().Exec()
	if err != nil {
		return nil, err
	}
	item, err := GetById[ReturnItem](table, id, nil)
	if err != nil {
		return nil, err
	}
	return item, nil
}

func Exists(table string, p QueryProcessor) (bool, error) {
	var id string
	query := p.Process(db.From(table).Select(goqu.C("id")).Limit(1))
	s, _, _ := query.ToSQL()
	fmt.Println(s)
	found, err := query.ScanVal(&id)
	return found, err
}
