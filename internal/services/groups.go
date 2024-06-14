package services

import (
	DB "github.com/Yakov-Varnaev/ft/internal"
	"github.com/Yakov-Varnaev/ft/internal/models"
	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

type Groups struct{}

func (s *Groups) GetById(id uuid.UUID) (*models.Group, error) {
	return DB.GetById[models.Group](DB.GROUPS_TABLE, id)
}

func (s *Groups) Create(group *models.Group) (*models.Group, error) {
	group.ID = uuid.New()
	db := DB.GetDB()
	_, err := db.Insert(DB.GROUPS_TABLE).Rows(group).Executor().Exec()
	if err != nil {
		return nil, err
	}
	resGroup := models.Group{}

	_, err = s.GetById(group.ID)
	if err != nil {
		return nil, err
	}

	return &resGroup, nil
}

// TODO: add pagination
func (s *Groups) List() (*[]models.Group, error) {
	db := DB.GetDB()
	groups := make([]models.Group, 0)
	err := db.From(DB.GROUPS_TABLE).ScanStructs(&groups)
	if err != nil {
		return nil, err
	}
	return &groups, nil
}

func (s *Groups) Update(id uuid.UUID, group models.WriteGroup) (*models.Group, error) {
	db := DB.GetDB()
	_, err := db.Update(DB.GROUPS_TABLE).Where(goqu.C("id").Eq(id)).Set(group).Executor().Exec()
	if err != nil {
		return nil, err
	}
	return s.GetById(id)
}

func (s *Groups) Delete(id uuid.UUID) (uuid.UUID, error) {
	db := DB.GetDB()
	_, err := db.Delete(DB.GROUPS_TABLE).Where(goqu.C("id").Eq(id)).Executor().Exec()
	if err != nil {
		return id, err
	}
	return id, nil
}
