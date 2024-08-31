package service

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/Yakov-Varnaev/ft/internal/config"
	"github.com/Yakov-Varnaev/ft/internal/database"
	"github.com/Yakov-Varnaev/ft/internal/model"
	groupRepo "github.com/Yakov-Varnaev/ft/internal/repository/group"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
	"github.com/Yakov-Varnaev/ft/pkg/repository/utils"
)

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
	cfg, err := config.New()
	if err != nil {
		t.Fatal(err)
	}
	cfg.DB.Database = "test_group"
	db := database.New(cfg.DB)
	defer func() {
		db.MustExec("TRUNCATE TABLE groups CASCADE;")
		db.Close()
	}()

	repo := groupRepo.New(db)
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

func Test_service_List(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Fatal(err)
	}
	cfg.DB.Database = "test_group"
	db := database.New(cfg.DB)
	defer func() {
		db.MustExec("TRUNCATE TABLE groups CASCADE;")
		db.Close()
	}()

	repo := groupRepo.New(db)
	s := New(repo)

	groups := make([]*model.Group, 0)
	for i := range 3 {
		name := fmt.Sprintf("Group %d", i)
		g, err := s.Create(&model.GroupInfo{Name: name})
		if err != nil {
			t.Fatal(err)
		}
		groups = append(groups, g)
	}

	type args struct {
		pg pagination.Pagination
	}
	tests := []struct {
		name    string
		args    args
		want    *pagination.Page[*model.Group]
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "first page all posts",
			args: args{
				pagination.Pagination{Limit: 10},
			},
			want: &pagination.Page[*model.Group]{
				Total: 3,
				Data:  groups,
			},
			wantErr: false,
		},
		{
			name: "limit 1 offset 0",
			args: args{
				pagination.Pagination{Limit: 1},
			},
			want: &pagination.Page[*model.Group]{
				Total: 3,
				Data:  groups[:1],
			},
		},
		{
			name: "limit 1 offset 1",
			args: args{
				pagination.Pagination{Limit: 1, Offset: 1},
			},
			want: &pagination.Page[*model.Group]{
				Total: 3,
				Data:  groups[1:2],
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.List(tt.args.pg)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_Delete(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Fatal(err)
	}
	cfg.DB.Database = "test_group"
	db := database.New(cfg.DB)
	defer func() {
		db.MustExec("TRUNCATE TABLE groups CASCADE;")
		db.Close()
	}()

	repo := groupRepo.New(db)
	s := New(repo)

	group, err := repo.Create(&model.GroupInfo{Name: "Test Group"})
	if err != nil {
		t.Fatalf("Cannot create group: %v", err)
	}

	err = s.Delete(group.ID)
	if err != nil {
		t.Fatalf("Cannot delete group: %v", err)
	}

	exists, err := repo.Exists(utils.Filters{"name": group.Name})
	if err != nil {
		t.Fatalf("Cannot check if post exists: %v", err)
	}
	if exists {
		t.Fatalf("Group exists after deletion.")
	}
}
