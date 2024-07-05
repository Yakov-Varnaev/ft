package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/Yakov-Varnaev/ft/internal/config"
	"github.com/Yakov-Varnaev/ft/internal/database"
	repository "github.com/Yakov-Varnaev/ft/internal/repository/group"
	repoModel "github.com/Yakov-Varnaev/ft/internal/repository/group/model"
	"github.com/Yakov-Varnaev/ft/internal/server"
	"github.com/Yakov-Varnaev/ft/internal/service/group/model"
	"github.com/Yakov-Varnaev/ft/pkg/repository/utils"
	"github.com/gavv/httpexpect/v2"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/suite"
)

type GroupHandlerSuite struct {
	suite.Suite
	pgdb   *sqlx.DB
	db     *sqlx.DB
	config config.Config
	server *server.Server
	groups []*model.Group
	repo   *repository.Repository
}

func (suite *GroupHandlerSuite) createGroups() {
	suite.groups = []*model.Group{}
	for i := range 3 {
		createdGroup, err := suite.repo.Create(
			&repoModel.GroupInfo{Name: fmt.Sprintf("Test Group %02d", i)},
		)
		suite.Require().NoError(err)
		group := model.FromRepoGroup(createdGroup)
		suite.groups = append(suite.groups, group)
	}
}

func (suite *GroupHandlerSuite) SetupSuite() {
	config, err := config.New()
	suite.Require().NoError(err)
	config.Database.Database = "postgres"
	suite.pgdb = database.New(config.Database)
	suite.pgdb.MustExec("CREATE DATABASE test_group")

	config.Database.Database = "test_group"

	sqlPath := "./../create_tables.sql"
	c, err := os.ReadFile(sqlPath)
	suite.Require().NoError(err, "Cannot read sql file.")
	sql := string(c)

	db := database.New(config.Database)
	suite.db = db
	db.MustExec(sql)

	suite.repo = repository.New(db)
	suite.config = config
	suite.server = server.New(suite.config)
}

func (suite *GroupHandlerSuite) TearDownSuite() {
	fmt.Println("Tear Down suite")
	suite.db.Close()
	suite.server.Close()
	suite.pgdb.MustExec("DROP DATABASE test_group")
}

func (suite *GroupHandlerSuite) SetupTest() {
	suite.config.Database.Database = "test_group"
	suite.createGroups()
}

func (suite *GroupHandlerSuite) TearDownTest() {
	suite.db.Exec("TRUNCATE TABLE groups CASCADE")
}

func (s *GroupHandlerSuite) TestGroupsList() {
	serv := httptest.NewServer(s.server.Engine)
	defer serv.Close()
	e := httpexpect.Default(s.T(), serv.URL)

	obj := e.GET("/api/v1/groups/").
		Expect().
		Status(http.StatusOK).
		JSON().
		Object()
	groups := []interface{}{}
	for _, g := range s.groups {
		groups = append(groups, g)
	}
	obj.
		ContainsKey("total").
		HasValue("total", 3).
		ContainsKey("data").
		Value("data").
		Array().
		ConsistsOf(groups...)
}

func (s *GroupHandlerSuite) TestGroupCreate() {
	serv := httptest.NewServer(s.server.Engine)
	defer serv.Close()
	e := httpexpect.Default(s.T(), serv.URL)
	data := model.GroupInfo{Name: "new group"}

	e.POST("/api/v1/groups/").
		WithJSON(data).
		Expect().
		Status(http.StatusCreated).
		JSON().
		Object().
		ContainsKey("id").
		ContainsKey("name").
		HasValue("name", data.Name)

	exists, err := s.repo.Exists(utils.Filters{"name": "new group"})
	s.Require().NoError(err)
	s.Require().True(exists)
}

func (s *GroupHandlerSuite) TestGroupUpdate() {
	serv := httptest.NewServer(s.server.Engine)
	defer serv.Close()
	e := httpexpect.Default(s.T(), serv.URL)
	group := s.groups[0]
	data := model.GroupInfo{Name: "new group name"}

	e.PUT(fmt.Sprintf("/api/v1/groups/%s/", group.UUID)).
		WithJSON(data).
		Expect().
		Status(http.StatusOK).
		JSON().
		Object().
		HasValue("name", data.Name).
		HasValue("id", group.UUID)

	updatedGroup, err := s.repo.GetById(group.UUID)
	s.Require().NoError(err)
	s.Require().Equal(updatedGroup.Name, data.Name)
}
func (s *GroupHandlerSuite) TestGroupDelete() {
	serv := httptest.NewServer(s.server.Engine)
	defer serv.Close()
	e := httpexpect.Default(s.T(), serv.URL)

	group := s.groups[0]

	e.DELETE(fmt.Sprintf("/api/v1/groups/%s/", group.UUID)).
		Expect().
		Status(http.StatusNoContent)

	exists, err := s.repo.Exists(utils.Filters{"id": group.UUID})
	s.Require().NoError(err)
	s.Require().False(exists)
}

func TestGroupHandlerSuite(t *testing.T) {
	suite.Run(t, new(GroupHandlerSuite))
}
