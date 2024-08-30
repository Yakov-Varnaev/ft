package service

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/Yakov-Varnaev/ft/internal/config"
	"github.com/Yakov-Varnaev/ft/internal/database"
	"github.com/Yakov-Varnaev/ft/internal/model"
	repository "github.com/Yakov-Varnaev/ft/internal/repository/group"
)

// test create valid data
// test create invalid data(duplicating name)
// test list pagination

func TestMain(t *testing.M) {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	db := database.New(cfg.DB)
	db.MustExec("CREATE DATABASE test_group")
	db.Close()

	sqlPath := "./../../../create_tables.sql"
	c, err := os.ReadFile(sqlPath)
	if err != nil {
		panic(err)
	}
	sqlStmt := string(c)
	cfg.DB.Database = "test_group"
	db = database.New(cfg.DB)
	db.MustExec(sqlStmt)
	db.Close()

	fmt.Println("Main start")
	t.Run()
	fmt.Println("Clean up")

	cfg.DB.Database = "ft"
	db = database.New(cfg.DB)
	db.MustExec("DROP DATABASE test_group")
	db.Close()
}

func Test_service_Create_invalid_data(t *testing.T) {
	fmt.Println("invalid data success")
	cfg, err := config.New()
	if err != nil {
		t.Fatal(err)
	}
	cfg.DB.Database = "test_group"
	db := database.New(cfg.DB)
	defer db.Close()

	repo := repository.New(db)
	s := New(repo)

	_, err = repo.Create(&model.GroupInfo{Name: "Test Group"})
	if err != nil {
		t.Fatalf("Cannot create group: %v", err)
	}

	type args struct {
		info *model.GroupInfo
	}
	tests := []struct {
		name    string
		args    args
		want    *model.Group
		wantErr bool
	}{
		{
			name: "Empty name",
			args: args{
				info: &model.GroupInfo{Name: ""},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Duplicate name",
			args: args{
				info: &model.GroupInfo{Name: "Test Group"},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Create(tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.Create() = %v, want %v", got, tt.want)
				return
			}
			res := db.QueryRowx("select count(*) from groups")
			var cnt int
			if err := res.Scan(&cnt); err != nil {
				t.Errorf("Cannot count groups: %v", err)
				return
			}
			if cnt > 1 {
				t.Errorf("Group was created event though error was returned")
			}
		})
	}
}
