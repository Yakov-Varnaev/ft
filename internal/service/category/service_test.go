package service

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/Yakov-Varnaev/ft/internal/config"
	"github.com/Yakov-Varnaev/ft/internal/database"
	"github.com/Yakov-Varnaev/ft/internal/model"
	categoryRepo "github.com/Yakov-Varnaev/ft/internal/repository/category"
	groupRepo "github.com/Yakov-Varnaev/ft/internal/repository/group"
	"github.com/Yakov-Varnaev/ft/pkg/pagination"
)

func TestMain(t *testing.M) {
	cfg, err := config.New()
	if err != nil {
		panic(err)
	}
	db := database.New(cfg.DB)
	db.MustExec("CREATE DATABASE test_category")
	db.Close()

	sqlPath := "./../../../create_tables.sql"
	c, err := os.ReadFile(sqlPath)
	if err != nil {
		panic(err)
	}
	sqlStmt := string(c)
	cfg.DB.Database = "test_category"
	db = database.New(cfg.DB)
	db.MustExec(sqlStmt)
	db.Close()

	fmt.Println("Main start")
	t.Run()
	fmt.Println("Clean up")

	cfg.DB.Database = "ft"
	db = database.New(cfg.DB)
	db.MustExec("DROP DATABASE test_category")
	db.Close()
}

func Test_service_Create(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Fatal(err)
	}
	cfg.DB.Database = "test_category"
	db := database.New(cfg.DB)
	defer func() {
		db.MustExec("TRUNCATE TABLE groups CASCADE;")
		db.Close()
	}()

	gRepo := groupRepo.New(db)
	cRepo := categoryRepo.New(db)
	s := New(cRepo)

	group, err := gRepo.Create(&model.GroupInfo{Name: "Test Group"})
	if err != nil {
		t.Fatalf("Cannot create group: %v", err)
	}

	info := &model.CategoryInfo{
		GroupID: group.ID,
		Name:    "Test Category",
	}

	category, err := s.Create(info)
	if err != nil {
		t.Fatalf("Cannot create category with valid data: %v", err)
	}
	if category.Name != info.Name {
		t.Fatalf("Category created with wrong name. Got: %v, want: %v", category.Name, info.Name)
	}
	if category.Group.ID != info.GroupID {
		t.Fatalf("Category created with wrong group. Got: %v, want: %v", category.Group.ID, info.GroupID)
	}
}

func Test_service_List(t *testing.T) {
	cfg, err := config.New()
	if err != nil {
		t.Fatal(err)
	}
	cfg.DB.Database = "test_category"
	db := database.New(cfg.DB)
	defer func() {
		db.MustExec("TRUNCATE TABLE groups CASCADE;")
		db.Close()
	}()

	gRepo := groupRepo.New(db)
	cRepo := categoryRepo.New(db)
	s := New(cRepo)

	group, err := gRepo.Create(&model.GroupInfo{Name: "Test Group"})
	if err != nil {
		t.Fatalf("Cannot create group: %v", err)
	}

	categories := make([]*model.Category, 0)
	for i := range 3 {
		name := fmt.Sprintf("Category %d", i)
		g, err := s.Create(&model.CategoryInfo{Name: name, GroupID: group.ID})
		if err != nil {
			t.Fatal(err)
		}
		categories = append(categories, g)
	}

	type args struct {
		pg pagination.Pagination
	}
	tests := []struct {
		name    string
		args    args
		want    *pagination.Page[*model.Category]
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "first page all posts",
			args: args{
				pagination.Pagination{Limit: 10},
			},
			want: &pagination.Page[*model.Category]{
				Total: 3,
				Data:  categories,
			},
			wantErr: false,
		},
		{
			name: "limit 1 offset 0",
			args: args{
				pagination.Pagination{Limit: 1},
			},
			want: &pagination.Page[*model.Category]{
				Total: 3,
				Data:  categories[:1],
			},
		},
		{
			name: "limit 1 offset 1",
			args: args{
				pagination.Pagination{Limit: 1, Offset: 1},
			},
			want: &pagination.Page[*model.Category]{
				Total: 3,
				Data:  categories[1:2],
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
