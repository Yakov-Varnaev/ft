package server_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	DB "github.com/Yakov-Varnaev/ft/internal"
	"github.com/Yakov-Varnaev/ft/internal/models"
	"github.com/Yakov-Varnaev/ft/internal/server"
	"github.com/Yakov-Varnaev/ft/internal/services"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

func createGroups(n int64) []models.Group {
	groups := make([]models.Group, n)
	for i := range n {
		g := models.Group{
			ID:   uuid.New(),
			Name: fmt.Sprintf("Group %d", i),
		}
		_, err := DB.Create[models.Group, models.Group](
			DB.GROUPS_TABLE, &g,
		)
		if err != nil {
			panic(err.Error())
		}
		groups[i] = g
	}
	return groups
}

func createCategories(n int64, groups []models.Group) []models.DetailedCategory {
	s := new(services.Categories)
	for i := range n {
		g := rand.Int63n(n)
		c := models.WriteCategory{
			Name:  fmt.Sprintf("Category %d", i),
			Group: groups[g].ID.String(),
		}
		_, err := s.Create(&c)
		if err != nil {
			panic(err.Error())
		}
	}
	categories, err := s.List()
	if err != nil {
		panic(err.Error())
	}
	return *categories
}

type ApiTestSuite struct {
	suite.Suite
	groups []models.Group
	group  models.Group

	categories []models.DetailedCategory
	category   models.DetailedCategory
}

func (s *ApiTestSuite) SetupSuite() {
	DB.Init()
	_, err := DB.GetDB().Exec(`DELETE FROM "groups"`)
	if err != nil {
		s.FailNowf("teardown is fucked setup: %s", err.Error())
	}
	_, err = DB.GetDB().Exec(`DELETE FROM "categories"`)
	if err != nil {
		s.FailNowf("teardown is fucked setup: %s", err.Error())
	}
}

func (s *ApiTestSuite) SetupTest() {
	s.groups = createGroups(10)
	s.group = s.groups[0]

	s.categories = createCategories(10, s.groups)
	s.category = s.categories[0]
}

func (s *ApiTestSuite) TearDownTest() {
	_, err := DB.GetDB().Exec(`DELETE FROM "groups"`)
	if err != nil {
		s.FailNowf("teardown is fucked: %s", err.Error())
	}
	_, err = DB.GetDB().Exec(`DELETE FROM "categories"`)
	if err != nil {
		s.FailNowf("teardown is fucked: %s", err.Error())
	}

}

func (s *ApiTestSuite) testRead(url string, data interface{}) {

	router := server.SetupRouter()
	w := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, url, nil)
	s.Require().NoError(err)

	expectedJson, _ := json.Marshal(data)
	router.ServeHTTP(w, request)

	s.Require().Equal(http.StatusOK, w.Code)
	s.Require().Equal(
		string(expectedJson),
		w.Body.String(),
	)
}

func (s *ApiTestSuite) TestGroupsRead() {
	s.testRead("/api/v1/groups/", s.groups)
}

func (s *ApiTestSuite) TestGroupCreate() {
	router := server.SetupRouter()
	w := httptest.NewRecorder()
	data := models.WriteGroup{
		Name: "New group",
	}
	jsonData, _ := json.Marshal(data)
	request, err := http.NewRequest(
		http.MethodPost,
		"/api/v1/groups/",
		strings.NewReader(string(jsonData)),
	)
	s.Require().NoError(err)

	var resGroup models.Group
	q := DB.GetDB().From(DB.GROUPS_TABLE).Where(goqu.C("name").Eq(data.Name))
	sq, _, _ := q.ToSQL()
	fmt.Println(sq)
	found, _ := q.ScanStruct(&resGroup)
	if !found {
		panic("fuck")
	}

	expectedJson, _ := json.Marshal(resGroup)
	router.ServeHTTP(w, request)

	s.Require().Equal(http.StatusOK, w.Code)
	s.Require().Equal(string(expectedJson), w.Body.String())
}

func (s *ApiTestSuite) TestCategoriesRead() {
	s.testRead("/api/v1/categories/", s.categories)
}

func TestApiSuite(t *testing.T) {
	suite.Run(t, &ApiTestSuite{})
}
